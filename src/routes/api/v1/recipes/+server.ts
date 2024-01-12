import { json, type RequestHandler } from '@sveltejs/kit';
import { parseParamAsNumber } from '../util';
import { db } from '$lib/data/db';
import { recipe, recipeIngredient } from '$src/schema';
import { and, eq, inArray } from 'drizzle-orm';

export const GET: RequestHandler = async ({ url, fetch }) => {
	const includeIngredients = url.searchParams.get('ingredients') === 'true';
	const limit = parseParamAsNumber(url.searchParams.get('limit'), 10);
	const offset = parseParamAsNumber(url.searchParams.get('offset'), 0);
	const unpublished = url.searchParams.get('unpublished') === 'true';
	const fromUser = url.searchParams.get('user_name');

	const andPredicates = [];

	if (!unpublished) {
		andPredicates.push(eq(recipe.published, true));
	}

	if (fromUser) {
		const response = await fetch(`/api/v1/user?name=${fromUser}`);
		if (response.status !== 200) {
			return json({ message: 'No such user' }, { status: 500 });
		}
		const fromUserId = (await response.json()).id;
		andPredicates.push(eq(recipe.ownerId, fromUserId));
	}

	let resultRecipes = await db
		.select()
		.from(recipe)
		.limit(limit)
		.offset(offset)
		.where(and(...andPredicates));

	if (includeIngredients) {
		const recipeIds = resultRecipes.map((recipe) => recipe.id);
		const recipeIngredients = await db
			.select()
			.from(recipeIngredient)
			.where(inArray(recipeIngredient.recipeId, recipeIds));
		resultRecipes = resultRecipes.map((recipe) => {
			return {
				...recipe,
				ingredients: recipeIngredients.filter((ingredient) => ingredient.recipeId === recipe.id)
			};
		});
	}

	if (!recipe) {
		return json({ message: 'Recipe not found' }, { status: 404 });
	}

	return json(recipe);
};
