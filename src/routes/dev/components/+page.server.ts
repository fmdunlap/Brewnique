import { getRecipes } from '$lib/data/recipe.js';

export const load = async () => {
	return {
		recipes: await getRecipes(30, 0)
	};
};
