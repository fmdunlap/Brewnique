import { json, type RequestHandler } from '@sveltejs/kit';

function parseParamAsNumber(param: string | null, def: number): number {
	if (!param) {
		return def;
	}

	const parsed = Number.parseInt(param);

	if (Number.isNaN(parsed)) {
		return def;
	}

	return parsed;
}

export const GET: RequestHandler = async ({ url, locals: { supabase } }) => {
	const includeIngredients = url.searchParams.get('ingredients') === 'true';
	const limit = parseParamAsNumber(url.searchParams.get('limit'), 10);
	const offset = parseParamAsNumber(url.searchParams.get('offset'), 0);

	const selectString = includeIngredients
		? '*, ingredients: recipe_ingredient(name, quantity, unit, type)'
		: '*';

	const { data: recipe, error } = await supabase
		.from('recipe')
		.select(selectString)
		.range(offset, offset + limit - 1);

	if (error) {
		return json({ message: error.message }, { status: 500 });
	}

	if (!recipe) {
		return json({ message: 'Recipe not found' }, { status: 404 });
	}

	return json(recipe);
};
