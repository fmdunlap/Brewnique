import { auth, facebookAuth } from '$lib/auth/lucia.js';
import { OAuthRequestError } from '@lucia-auth/oauth';

export const GET = async ({ url, cookies, locals, params, route, request, platform }) => {
	const storedState = cookies.get('facebook_oauth_state');
	const state = url.searchParams.get('state');
	const code = url.searchParams.get('code');

	console.log('PARAMS: ', JSON.stringify(params, null, 2));
	console.log('ROUTE: ', JSON.stringify(route, null, 2));
	console.log('REQUEST: ', JSON.stringify(request.headers, null, 2));
	console.log('COOKIES: ', JSON.stringify(cookies, null, 2));
	console.log('URL: ', JSON.stringify(url, null, 2));
	console.log('PLATFORM: ', JSON.stringify(platform, null, 2));

	// validate state
	if (!storedState || !state || storedState !== state || !code) {
		return new Response(null, {
			status: 400
		});
	}
	try {
		const { getExistingUser, facebookUser, createUser } = await facebookAuth.validateCallback(code);

		console.log(JSON.stringify(facebookUser, null, 2));

		const getUser = async () => {
			const existingUser = await getExistingUser();
			if (existingUser) return existingUser;
			const user = await createUser({
				attributes: {
					username: facebookUser.email!
				}
			});
			return user;
		};

		const user = await getUser();
		const session = await auth.createSession({
			userId: user.userId,
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
