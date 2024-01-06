import { getUserProfile } from '$lib/data/profile.js';

export async function load({ parent }) {
	const { supabase, session } = await parent();

	return {
		supabase,
		session,
		user_profile: session ? await getUserProfile(session, supabase) : null
	};
}
