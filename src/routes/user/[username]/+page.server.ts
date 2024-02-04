import { getUserProfileByUsername } from '$lib/data/user.js';
import { getAllRecipesByUser, getPublishedRecipesByUser } from '$lib/data/recipe.js';
import { fail } from '@sveltejs/kit';

export const load = async ({ params, locals }) => {
	const username = params.username;
	if (!username) {
		return fail(400, { message: 'Bad request' });
	}

	const user = await getUserProfileByUsername(username);
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

export const actions = {
	newAvatar: async ({ params, locals, request }) => {
		const session = await locals.auth.validate();
		if (!session || !session.user) {
			return fail(401, { message: 'You must be logged in to edit a recipe' });
		}

		console.log('Got avatar');

		const form = await request.formData();
		const avatar = form.get('avatar');
		if (!avatar) {
			return fail(400, { message: 'Bad request' });
		}
	}
};
