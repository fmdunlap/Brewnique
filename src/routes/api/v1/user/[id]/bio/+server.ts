import { db } from '$lib/data/db';
import { obscenityMatcher } from '$lib/utils';
import { user } from '$src/schema';
import { type RequestHandler } from '@sveltejs/kit';
import { eq } from 'drizzle-orm';

export const POST: RequestHandler = async ({ request, params, locals }) => {
	const userId = params.id;

	if (!userId) {
		return new Response('User ID not provided', { status: 400 });
	}

	const validatedId = (await locals.auth.validate())?.user.userId;

	if (!validatedId || userId != validatedId) {
		return new Response('Unauthorized', { status: 401 });
	}

	const newBio = (await request.json()).bio;

	if (newBio.length > 160) {
		return new Response('Bio too long', { status: 400 });
	}

	if (obscenityMatcher.hasMatch(newBio)) {
		return new Response('Bio contains profanity', { status: 400 });
	}

	await db.update(user).set({ bio: newBio }).where(eq(user.id, userId));
	return new Response('Bio updated', { status: 200 });
};
