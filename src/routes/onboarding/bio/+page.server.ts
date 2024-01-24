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

		const form = await superValidate(request, bioFormSchema);
		if (!form.valid) {
			return fail(400, { form });
		}

		// Skip the bio if the user wants to
		if (form.data.skip) {
			await advanceOnboardingState(session.user.userId);
			return { form };
		}

		// Check for profanity in the bio
		if (obscenityMatcher.hasMatch(form.data.bio)) {
			form.errors.bio = ['Bio contains profanity.'];
			return fail(400, { form });
		}

		await setBio(session.user.userId, form.data.bio);
		await advanceOnboardingState(session.user.userId);
		auth.updateUserAttributes(session.user.userId, {
			bio: form.data.bio,
			onboarding_status: 'COMPLETE'
		});

		await auth.validateSession(session.sessionId);

		throw redirect(302, '/');
	}
};
