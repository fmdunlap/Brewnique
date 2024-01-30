import type { Actions } from '@sveltejs/kit';
import { getRecipes } from '$lib/data/recipe';

export const load = async () => {
	return {
		recipes: await getRecipes(30, 0, true, null)
	};
};

export const actions = {
	default: async ({ request }) => {}
} satisfies Actions;
