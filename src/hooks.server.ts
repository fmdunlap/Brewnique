import type { Handle } from '@sveltejs/kit';
// import { AUTH_SECRET, GITHUB_ID, GITHUB_SECRET } from '$env/static/private';

import { auth } from '$lib/auth/lucia';

export const handle: Handle = async ({ event, resolve }) => {
	// we can pass `event` because we used the SvelteKit middleware
	event.locals.auth = auth.handleRequest(event);
	return await resolve(event);
};
