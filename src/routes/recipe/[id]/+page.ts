import { error } from '@sveltejs/kit';

export const load = async ({ params, fetch }) => {
	const recipe_resp = await fetch(`/api/v1/recipe?id=${params.id}&ingredients=true`);
	if (recipe_resp.status !== 200) {
		console.log(await recipe_resp.json());
		throw error(404, 'Recipe not found');
	}
	const recipe = await recipe_resp.json();

	const user_resp = await fetch(`/api/v1/user?id=${recipe.user_id}`);
	const recipe_owner = await user_resp.json();

	return {
		recipe,
		recipe_owner
	};
};
