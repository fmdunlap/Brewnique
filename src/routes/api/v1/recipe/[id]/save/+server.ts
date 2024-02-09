import { saveRecipe } from '$lib/data/recipe';
import type { RequestHandler } from '@sveltejs/kit';

export const POST: RequestHandler = async ({ params, locals }) => {
	const recipeId = params.id;
	if (!recipeId) {
		return new Response('No recipe specified', { status: 404 });
	}

	const session = await locals.auth.validate();
	if (!session) {
		return new Response('Session is not valid', { status: 401 });
	}

	await saveRecipe(recipeId, session.user.userId);

	return new Response('saved recipe', { status: 200 });
};
