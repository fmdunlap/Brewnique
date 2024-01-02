import type { EmailOtpType } from '@supabase/supabase-js';
import { redirect } from '@sveltejs/kit';

export const GET = async (event) => {
	const {
		url,
		locals: { supabase }
	} = event;
	const token_hash = url.searchParams.get('token_hash') as string;
	const token_type = url.searchParams.get('type');
	const next = url.searchParams.get('next') ?? '/';

	if (
		token_hash &&
		token_type &&
		['signup', 'invite', 'magiclink', 'recovery', 'email_change', 'email'].includes(token_type)
	) {
		const { error } = await supabase.auth.verifyOtp({
			token_hash,
			type: token_type as EmailOtpType
		});
		if (!error) {
			throw redirect(303, `/${next.slice(1)}`);
		}
	}

	throw redirect(303, '/auth/auth-code-error');
};
