import { addDefaultAvatarToStorage, userHasAvatar } from '$lib/data/avatar.js';
import { redirect } from '@sveltejs/kit';

export const GET = async (event) => {
	const {
		url,
		locals: { supabase }
	} = event;

	const code = url.searchParams.get('code') as string;
	const next = url.searchParams.get('next') ?? '/';
	if (code) {
		const {
			error,
			data: { user }
		} = await supabase.auth.exchangeCodeForSession(code);
		if (!error) {
			if (!user?.id) {
				throw redirect(303, '/login/error');
			}
			if (!(await userHasAvatar(user.id, supabase))) {
				addDefaultAvatarToStorage(user.id, supabase);
			}
			throw redirect(303, `/${next.slice(1)}`);
		}
	}

	throw redirect(303, '/login/error');
};
