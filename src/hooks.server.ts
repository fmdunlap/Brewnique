import { SvelteKitAuth } from '@auth/sveltekit';
import { DrizzleAdapter } from '@auth/drizzle-adapter';
import GitHub from '@auth/sveltekit/providers/github';
import type { Handle } from '@sveltejs/kit';
import { AUTH_SECRET, GITHUB_ID, GITHUB_SECRET } from '$env/static/private';
import { db } from '$lib/data/db';

export const handle = SvelteKitAuth(async () => {
	const authOptions = {
		adapter: DrizzleAdapter(db),
		providers: [
			GitHub({
				clientId: GITHUB_ID,
				clientSecret: GITHUB_SECRET
			})
		],
		secret: AUTH_SECRET,
		trustHost: true
	};
	return authOptions;
}) satisfies Handle;
