import { getRecipeThreads, getRecipeWithIngredients } from '$lib/data/recipe.js';
import { getUserProfileById } from '$lib/data/user.js';
import { error } from '@sveltejs/kit';

export const load = async ({ params }) => {
	const recipeId = params.id;

	if (!recipeId) {
		return error(400, { message: 'Bad request' });
	}

	const recipe = await getRecipeWithIngredients(recipeId);

	if (!recipe) {
		return error(404, { message: 'Recipe not found' });
	}

	return {
		recipe,
		recipeOwner: await getUserProfileById(recipe.ownerId),
		threads: await getRecipeThreads(recipeId)
	};
};
