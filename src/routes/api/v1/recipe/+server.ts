import { json, type RequestHandler } from '@sveltejs/kit';

export const GET: RequestHandler = async ({ url }) => {
	const recipeId = url.searchParams.get('id');
	const includeIngredients = url.searchParams.get('ingredients') === 'true';

	if (!recipeId) {
		return json({ message: 'Missing recipe id' }, { status: 400 });
	}

	const selectString = includeIngredients
		? '*, ingredients: recipe_ingredient(name, quantity, unit, type)'
		: '*';

	const { data: recipe, error } = await supabase
		.from('recipe')
		.select(selectString)
		.eq('id', recipeId)
		.single();

	if (error) {
		return json({ message: error.message }, { status: 500 });
	}

	if (!recipe) {
		return json({ message: 'Recipe not found' }, { status: 404 });
	}

	return json(recipe);
};
