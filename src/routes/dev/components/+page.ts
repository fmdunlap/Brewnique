import { error } from '@sveltejs/kit';

export const load = async ({ fetch }) => {
	const recipes_resp = await fetch(`/api/v1/recipes`);
	if (recipes_resp.status !== 200) {
		console.log(await recipes_resp.json());
		throw error(404, 'Recipe not found');
	}
	const recipes = await recipes_resp.json();
	return {
		recipes
	};
};
