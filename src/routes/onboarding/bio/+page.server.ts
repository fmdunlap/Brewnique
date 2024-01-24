import { fail, redirect } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms/server';
import { bioFormSchema } from './BioForm.js';
import { advanceOnboardingState, obscenityMatcher } from '../utils.js';
import { db } from '$lib/data/db.js';
import { user } from '$src/schema.js';
import { eq } from 'drizzle-orm';
import { auth } from '$lib/auth/lucia';

export const load = async () => {
	return {
		form: await superValidate(bioFormSchema)
	};
};

async function setBio(userId: string, bio: string) {
	return await db.update(user).set({ bio: bio }).where(eq(user.id, userId));
}

export const actions = {
	default: async ({ request, locals }) => {
		const session = await locals.auth.validate();
		if (!session) {
			return fail(401, { message: 'Unauthorized' });
		}

		const bioForm = await superValidate(request, bioFormSchema);
		if (!bioForm.valid) {
			return fail(400, { bioForm });
		}

		// Skip the bio if the user wants to
		if (bioForm.data.skip) {
			await advanceOnboardingState(session.user.userId);
			return { bioForm };
		}

		// Check for profanity in the bio
		if (obscenityMatcher.hasMatch(bioForm.data.bio)) {
			bioForm.errors.bio = ['Bio contains profanity.'];
			return fail(400, { bioForm });
		}

		await setBio(session.user.userId, bioForm.data.bio);
		await advanceOnboardingState(session.user.userId);
		auth.updateUserAttributes(session.user.userId, {
			bio: bioForm.data.bio,
			onboarding_status: 'COMPLETE'
		});

		await auth.validateSession(session.sessionId);

		throw redirect(302, '/');
	}
};
