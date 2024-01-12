import { superValidate } from 'sveltekit-superforms/server';
import { bioFormSchema, displayNameFormSchema } from '$lib/types/forms';
import { fail } from '@sveltejs/kit';
import { RegExpMatcher, englishDataset, englishRecommendedTransformers } from 'obscenity';
import { db } from '$lib/data/db.js';
import { user } from '$src/schema.js';
import { eq } from 'drizzle-orm';
import { auth } from '$lib/auth/lucia.js';
import { goto } from '$app/navigation';

const obscenityMatcher = new RegExpMatcher({
	...englishDataset.build(),
	...englishRecommendedTransformers
});

export const load = async () => {
	return {
		displayNameForm: await superValidate(displayNameFormSchema),
		bioForm: await superValidate(bioFormSchema)
	};
};

async function usernameTaken(username: string) {
	return (await db.select().from(user).where(eq(user.username, username))).length > 0;
}

async function setUsername(userId: string, username: string) {
	return await db.update(user).set({ username: username }).where(eq(user.id, userId));
}

async function setBio(userId: string, bio: string) {
	return await db.update(user).set({ bio: bio }).where(eq(user.id, userId));
}

async function advanceOnboardingState(userId: string) {
	const currentState = (await db.select().from(user).where(eq(user.id, userId)))[0]
		.onboardingStatus;

	// The state machine is just a linear progression.
	//
	// display_name_pending -> bio_pending -> avatar_pending -> completed
	let nextState: typeof currentState | null = null;
	switch (currentState) {
		case 'PENDING_USERNAME':
			nextState = 'PENDING_BIO';
			break;
		// For now we skip the avatar upload
		case 'PENDING_BIO':
		case 'PENDING_AVATAR':
			nextState = 'COMPLETE';
			break;
		case 'COMPLETE':
		default:
			nextState = null;
	}
	if (!nextState) {
		return;
	}

	return await db.update(user).set({ onboardingStatus: nextState }).where(eq(user.id, userId));
}

export const actions = {
	display_name: async ({ request, locals }) => {
		const session = await locals.auth.validate();
		if (!session) {
			return fail(401, { message: 'Unauthorized' });
		}
		const displayNameForm = await superValidate(request, displayNameFormSchema);

		if (!displayNameForm.valid) {
			return fail(400, { displayNameForm });
		}

		// Check for duplicate names
		if (await usernameTaken(displayNameForm.data.display_name)) {
			displayNameForm.errors.display_name = ['Display name is already taken.'];
			return fail(400, { displayNameForm });
		}

		// Check for profanity in the display name
		if (obscenityMatcher.hasMatch(displayNameForm.data.display_name)) {
			displayNameForm.errors.display_name = ['Display name contains profanity.'];
			return fail(400, { displayNameForm });
		}

		await setUsername(session.user.userId, displayNameForm.data.display_name);
		await advanceOnboardingState(session.user.userId);
		auth.invalidateSession(session.id);

		return { displayNameForm };
	},
	bio: async ({ request, locals }) => {
		const session = await locals.auth.validate();
		if (!session) {
			return fail(401, { message: 'Unauthorized' });
		}

		const bioForm = await superValidate(request, bioFormSchema);
		console.log(JSON.stringify(bioForm));
		if (!bioForm.valid) {
			return fail(400, { bioForm });
		}

		// Skip the bio if the user wants to
		if (bioForm.data.skip) {
			console.log('skipping bio form');

			await advanceOnboardingState(session.user.userId);
			return { bioForm };
		}

		// Check for profanity in the bio
		if (obscenityMatcher.hasMatch(bioForm.data.bio)) {
			console.log('bio contains profanity');
			bioForm.errors.bio = ['Bio contains profanity.'];
			return fail(400, { bioForm });
		}

		await setBio(session.user.userId, bioForm.data.bio);
		await advanceOnboardingState(session.user.userId);
		goto('/');

		return { bioForm };
	}
};
