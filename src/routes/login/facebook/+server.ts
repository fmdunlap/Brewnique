import { dev } from '$app/environment';
import { facebookAuth } from '$lib/auth/lucia.js';

export const GET = async ({ cookies, url }) => {
	const [authUrl, state] = await facebookAuth(url.origin).getAuthorizationUrl();
	// store state
	cookies.set('facebook_oauth_state', state, {
		httpOnly: true,
		secure: !dev,
		path: '/',
		maxAge: 60 * 60
	});
	return new Response(null, {
		status: 302,
		headers: {
			Location: authUrl.toString()
		}
	});
};
