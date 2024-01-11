import { fail, redirect } from '@sveltejs/kit';
import { auth } from '$lib/auth/lucia';

export const GET = async ({ locals }) => {
	const session = await locals.auth.validate();
	if (!session) throw fail(401);
	await auth.invalidateSession(session.sessionId); // invalidate session
	locals.auth.setSession(null); // remove cookie
	throw redirect(302, '/'); // redirect to login page
};
