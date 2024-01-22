import { auth, githubAuth, getOrCreateOAuthUser } from '$lib/auth/lucia.js';
import { OAuthRequestError } from '@lucia-auth/oauth';

export const GET = async ({ url, cookies, locals }) => {
	const storedState = cookies.get('github_oauth_state');
	const state = url.searchParams.get('state');
	const code = url.searchParams.get('code');

	// validate state
	if (!storedState || !state || storedState !== state || !code) {
		return new Response(null, {
			status: 400
		});
	}
	try {
		const { githubUser } = await githubAuth(url.origin).validateCallback(code);

		if (!githubUser.email) {
			return new Response(null, {
				status: 400
			});
		}

		const user = await getOrCreateOAuthUser(githubUser.email, 'github', githubUser.id.toString());
		const session = await auth.createSession({
			userId: user.id,
			attributes: {}
		});
		locals.auth.setSession(session);
		return new Response(null, {
			status: 302,
			headers: {
				Location: '/'
			}
		});
	} catch (e) {
		if (e instanceof OAuthRequestError) {
			// invalid code
			return new Response(null, {
				status: 400
			});
		}
		console.log(JSON.stringify(e));
		return new Response(null, {
			status: 500
		});
	}
};
