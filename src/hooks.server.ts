import type { Handle } from '@sveltejs/kit';

import { auth } from '$lib/auth/lucia';

export const handle: Handle = async ({ event, resolve }) => {
	// we can pass `event` because we used the SvelteKit middleware
	event.locals.auth = auth.handleRequest(event);
	return await resolve(event);
};
