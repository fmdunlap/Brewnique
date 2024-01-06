import { superValidate } from 'sveltekit-superforms/server';
import { bioFormSchema, displayNameFormSchema } from '$lib/types/forms';
import { fail } from '@sveltejs/kit';
import type { SupabaseClient } from '@supabase/supabase-js';
import type { Database } from '$lib/types/supabaseDB.js';
import { RegExpMatcher, englishDataset, englishRecommendedTransformers } from 'obscenity';

const obscenity_matcher = new RegExpMatcher({
	...englishDataset.build(),
	...englishRecommendedTransformers
});

export const load = async () => {
	return {
		displayNameForm: await superValidate(displayNameFormSchema),
		bioForm: await superValidate(bioFormSchema)
	};
};

function displayNameTaken(supabase: SupabaseClient<Database>, display_name: string) {
	return supabase
		.from('profile')
		.select('id')
		.eq('display_name', display_name)
		.then(({ data }) => {
			console.log(JSON.stringify(data));
			return data != null && data.length > 0;
		});
}

async function setDisplayName(
	supabase: SupabaseClient<Database>,
	user_id: string,
	display_name: string
) {
	return await supabase.from('profile').update({ display_name }).eq('id', user_id);
}

async function advanceOnboardingState(supabase: SupabaseClient<Database>, user_id: string) {
	const current_state = (
		await supabase.from('profile').select('onboarding_state').eq('id', user_id).single()
	).data?.onboarding_state;

	// The state machine is just a linear progression.
	//
	// email_unconfirmed -> display_name_pending -> bio_pending -> avatar_pending -> completed
	let next_state: Database['public']['Enums']['onboarding_state'] | null = null;
	switch (current_state) {
		case 'email_unconfirmed':
			next_state = 'display_name_pending';
			break;
		case 'display_name_pending':
			next_state = 'bio_pending';
			break;
		// For now we skip the avatar upload
		case 'bio_pending':
		case 'avatar_pending':
			next_state = 'completed';
			break;
		case 'completed':
		default:
			next_state = null;
	}
	if (!next_state) {
		return;
	}

	return await supabase.from('profile').update({ onboarding_state: next_state }).eq('id', user_id);
}

export const actions = {
	display_name: async ({ request, locals }) => {
		const session = await locals.getSession();
		if (!session) {
			return fail(401, { message: 'Unauthorized' });
		}
		const displayNameForm = await superValidate(request, displayNameFormSchema);

		if (!displayNameForm.valid) {
			return fail(400, { displayNameForm });
		}

		// Check for duplicate names
		if (await displayNameTaken(locals.supabase, displayNameForm.data.display_name)) {
			displayNameForm.errors.display_name = ['Display name is already taken.'];
			return fail(400, { displayNameForm });
		}

		// Check for profanity in the display name
		if (obscenity_matcher.hasMatch(displayNameForm.data.display_name)) {
			displayNameForm.errors.display_name = ['Display name contains profanity.'];
			return fail(400, { displayNameForm });
		}

		await setDisplayName(locals.supabase, session.user.id, displayNameForm.data.display_name);
		await advanceOnboardingState(locals.supabase, session.user.id);

		return { displayNameForm };
	},
	bio: async ({ request, locals }) => {
		const session = await locals.getSession();
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
			await advanceOnboardingState(locals.supabase, session.user.id);
			return { bioForm };
		}

		// Check for profanity in the bio
		if (obscenity_matcher.hasMatch(bioForm.data.bio)) {
			bioForm.errors.bio = ['Bio contains profanity.'];
			return fail(400, { bioForm });
		}

		return { bioForm };
	}
};
