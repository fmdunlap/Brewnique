import { auth, facebookAuth, getOrCreateOAuthUser } from '$lib/auth/lucia.js';
import { OAuthRequestError } from '@lucia-auth/oauth';

export const GET = async ({ url, cookies, locals }) => {
	const storedState = cookies.get('facebook_oauth_state');
	const state = url.searchParams.get('state');
	const code = url.searchParams.get('code');

	// validate state
	if (!storedState || !state || storedState !== state || !code) {
		return new Response(null, {
			status: 400
		});
	}
	try {
		const { facebookUser } = await facebookAuth(url.origin).validateCallback(code);

		if (!facebookUser.email) {
			return new Response(null, {
				status: 400
			});
		}
		const user = await getOrCreateOAuthUser(facebookUser.email, 'facebook', facebookUser.id);
		const session = await auth.createSession({
			userId: user.id,
			attributes: {}
		});
		locals.auth.setSession(session);
		if (session.user.onboardingStatus !== 'COMPLETE') {
			return new Response(null, {
				status: 302,
				headers: {
					Location: '/onboarding'
				}
			});
		}
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
