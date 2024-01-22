import { getUserProfile } from '$lib/data/user.js';
import { getAllRecipesByUser, getPublishedRecipesByUser } from '$lib/data/recipe.js';
import { fail } from '@sveltejs/kit';

export const load = async ({ params, locals }) => {
	const username = params.username;
	if (!username) {
		return fail(400, { message: 'Bad request' });
	}

	const user = await getUserProfile(username);
	if (!user) {
		return fail(404, { message: 'User not found' });
	}

	const session = await locals.auth.validate();

	let recipes;
	if (session && user.id == session.user.userId) {
		recipes = await getAllRecipesByUser(user.id);
	} else {
		recipes = await getPublishedRecipesByUser(user.id);
	}

	return {
		user,
		recipes
	};
};
