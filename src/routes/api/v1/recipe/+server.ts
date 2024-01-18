import { db } from '$lib/data/db';
import { recipe, recipeIngredient } from '$src/schema';
import { json, type RequestHandler } from '@sveltejs/kit';
import { eq } from 'drizzle-orm';

// TODO: I can probably make this more efficient by using a join.
export const GET: RequestHandler = async ({ url }) => {
	const recipeId = url.searchParams.get('id')!;
	const includeIngredients = url.searchParams.get('ingredients') === 'true';

	if (!recipeId) {
		return json({ message: 'Missing recipe id' }, { status: 400 });
	}

	const recipeResult = await db.select().from(recipe).where(eq(recipe.id, recipeId));

	console.log('Recipe Result: ', JSON.stringify(recipeResult, null, 2));

	if (recipeResult.length < 1) {
		return json({ message: 'Recipe not found' }, { status: 404 });
	}

	if (!includeIngredients) {
		return json({ ...recipeResult[0] });
	}

	const recipeIngredients = await db
		.select()
		.from(recipeIngredient)
		.where(eq(recipeIngredient.recipeId, recipeResult[0].id));
	return json({ ...recipeResult[0], ingredients: recipeIngredients });
};
