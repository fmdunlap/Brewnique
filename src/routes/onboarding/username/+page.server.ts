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
		userForm: await superValidate(usernameFormSchema)
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

		const userForm = await superValidate(request, usernameFormSchema);
		console.log('POST', userForm);

		if (!userForm.valid) {
			console.log('Form invalid');
			return fail(400, { userForm });
		}

		// Check for duplicate names
		if (await usernameTaken(userForm.data.username)) {
			userForm.errors.username = ['Username is already taken.'];
			return fail(400, { userForm });
		}

		// Check for profanity in the username
		if (obscenityMatcher.hasMatch(userForm.data.username)) {
			userForm.errors.username = ['Username name contains profanity.'];
			return fail(400, { userForm });
		}

		await setUsername(session.user.userId, userForm.data.username);
		await advanceOnboardingState(session.user.userId);
		auth.updateUserAttributes(session.user.userId, {
			username: userForm.data.username,
			onboarding_status: 'PENDING_BIO'
		});

		throw redirect(302, '/onboarding/bio');
	}
};
