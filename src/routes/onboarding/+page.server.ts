import { superValidate } from 'sveltekit-superforms/server';
import { displayNameFormSchema } from '$lib/types/forms';
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
		form: await superValidate(displayNameFormSchema)
	};
};

function displayNameTaken(supabase: SupabaseClient<Database>, display_name: string) {
	// This is just an example. You should probably use a stored procedure.
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

	let next_state: Database['public']['Enums']['onboarding_state'] | null = null;
	switch (current_state) {
		case 'email_unconfirmed':
			next_state = 'display_name_pending';
			break;
		case 'display_name_pending':
			next_state = 'bio_pending';
			break;
		case 'bio_pending':
			next_state = 'avatar_pending';
			break;
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
		const form = await superValidate(request, displayNameFormSchema);

		// Convenient validation check:
		if (!form.valid) {
			// Again, return { form } and things will just work.
			return fail(400, { form });
		}

		// Check for duplicate names
		if (await displayNameTaken(locals.supabase, form.data.display_name)) {
			form.errors.display_name = ['Display name is already taken.'];
			return fail(400, { form });
		}

		// Check for profanity in the display name
		if (obscenity_matcher.hasMatch(form.data.display_name)) {
			form.errors.display_name = ['Display name contains profanity.'];
			return fail(400, { form });
		}

		const { error: dn_error } = await setDisplayName(
			locals.supabase,
			session.user.id,
			form.data.display_name
		);
		if (dn_error) {
			console.log(dn_error);
		}
		await advanceOnboardingState(locals.supabase, session.user.id);

		// Yep, return { form } here too
		return { form };
	}
};
