import { db } from '$lib/data/db.js';
import { user } from '$src/schema.js';
import { fail, redirect } from '@sveltejs/kit';
import { eq } from 'drizzle-orm';
import { superValidate } from 'sveltekit-superforms/server';
import { usernameFormSchema } from './UsernameForm.js';
import { auth } from '$lib/auth/lucia.js';
import { advanceOnboardingState, obscenityMatcher } from '../utils.js';

export const load = async () => {
	return {
		usernameForm: await superValidate(usernameFormSchema)
	};
};

async function usernameTaken(username: string) {
	return (await db.select().from(user).where(eq(user.username, username))).length > 0;
}

async function setUsername(userId: string, username: string) {
	return await db.update(user).set({ username: username }).where(eq(user.id, userId));
}

export const actions = {
	default: async ({ request, locals }) => {
		const session = await locals.auth.validate();
		if (!session) {
			return fail(401, { message: 'Unauthorized' });
		}
		const usernameForm = await superValidate(request, usernameFormSchema);

		if (!usernameForm.valid) {
			return fail(400, { usernameForm });
		}

		// Check for duplicate names
		if (await usernameTaken(usernameForm.data.username)) {
			usernameForm.errors.username = ['Username is already taken.'];
			return fail(400, { usernameForm });
		}

		// Check for profanity in the username
		if (obscenityMatcher.hasMatch(usernameForm.data.username)) {
			usernameForm.errors.username = ['Username name contains profanity.'];
			return fail(400, { usernameForm });
		}

		await setUsername(session.user.userId, usernameForm.data.username);
		await advanceOnboardingState(session.user.userId);
		auth.updateUserAttributes(session.user.userId, {
			username: usernameForm.data.username,
			onboarding_status: 'PENDING_BIO'
		});

		throw redirect(302, '/onboarding/bio');
	}
};
