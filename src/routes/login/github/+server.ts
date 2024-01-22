import { dev } from '$app/environment';
import { githubAuth } from '$lib/auth/lucia.js';

export const GET = async ({ cookies, url }) => {
	const [authUrl, state] = await githubAuth(url.origin).getAuthorizationUrl();
	// store state
	cookies.set('github_oauth_state', state, {
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
