import { getUserSavedRecipes } from '$lib/data/recipe';
import type { RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ locals, params }) => {
	const paramId = params.id;
	if (!paramId) return new Response('User not found', { status: 404 });

	const userSession = await locals.auth.validate();

	if (!userSession || paramId != userSession.user.userId)
		return new Response('Not authorized', {
			status: 401
		});

	return new Response(JSON.stringify(await getUserSavedRecipes(userSession.user.userId)), {
		status: 200
	});
};
