import { PUBLIC_SUPABASE_ANON_KEY, PUBLIC_SUPABASE_URL } from '$env/static/public';
import { createBrowserClient, isBrowser, parse } from '@supabase/ssr';
import type { LayoutLoad } from './$types';
import { type Database } from '$lib/types/supabaseDB';
import { getUserProfile } from '$lib/data/profile';

export const load: LayoutLoad = async ({ fetch, data, depends }) => {
	depends('supabase:auth');

	const supabase = createBrowserClient<Database>(PUBLIC_SUPABASE_URL, PUBLIC_SUPABASE_ANON_KEY, {
		global: {
			fetch
		},
		cookies: {
			get(key) {
				if (!isBrowser()) return JSON.stringify(data.session);

				const cookie = parse(document.cookie);
				return cookie[key];
			}
		}
	});

	const {
		data: { session }
	} = await supabase.auth.getSession();

	const user_profile = session ? await getUserProfile(session, supabase) : null;

	if (user_profile) {
		user_profile.data.avatar_url =
			'https://api.dicebear.com/7.x/lorelei-neutral/svg?' +
			session?.user.user_metadata.name.split(' ')[0] +
			'.svg';
	}

	return { supabase, session, user_profile };
};
