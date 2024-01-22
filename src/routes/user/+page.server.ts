import { redirect } from '@sveltejs/kit';

export const load = async ({ locals }) => {
	const session = await locals.auth.validate();
	if (!session || !session.user) {
		throw redirect(302, '/login');
	}
	throw redirect(302, '/user/' + session.user.username);
};
