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
		form: await superValidate(usernameFormSchema)
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
		const form = await superValidate(request, usernameFormSchema);

		if (!form.valid) {
			return fail(400, { form });
		}

		// Check for duplicate names
		if (await usernameTaken(form.data.username)) {
			form.errors.username = ['Username is already taken.'];
			return fail(400, { form });
		}

		// Check for profanity in the username
		if (obscenityMatcher.hasMatch(form.data.username)) {
			form.errors.username = ['Username name contains profanity.'];
			return fail(400, { form });
		}

		await setUsername(session.user.userId, form.data.username);
		await advanceOnboardingState(session.user.userId);
		auth.updateUserAttributes(session.user.userId, {
			username: form.data.username,
			onboarding_status: 'PENDING_BIO'
		});

		throw redirect(302, '/onboarding/bio');
	}
};
