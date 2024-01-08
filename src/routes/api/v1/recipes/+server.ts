import { json, type RequestHandler } from '@sveltejs/kit';
import { parseParamAsNumber } from '../util';

export const GET: RequestHandler = async ({ url, locals: { supabase }, fetch }) => {
	const includeIngredients = url.searchParams.get('ingredients') === 'true';
	const limit = parseParamAsNumber(url.searchParams.get('limit'), 10);
	const offset = parseParamAsNumber(url.searchParams.get('offset'), 0);
	const unpublished = url.searchParams.get('unpublished') === 'true';

	const fromUser = url.searchParams.get('user_name');

	let selectString = '*';

	if (includeIngredients) {
		selectString += ', ingredients: recipe_ingredient(name, quantity, unit, type)';
	}

	let query = supabase
		.from('recipe')
		.select(selectString)
		.range(offset, offset + limit - 1);

	if (!unpublished) {
		query = query.eq('published', true);
	}

	if (fromUser) {
		const response = await fetch(`/api/v1/user?name=${fromUser}`);
		if (response.status !== 200) {
			return json({ message: 'No such user' }, { status: 500 });
		}
		query = query.eq('user_id', (await response.json()).id);
	}

	const { data: recipe, error } = await query;

	if (error) {
		return json({ message: error.message }, { status: 500 });
	}

	if (!recipe) {
		return json({ message: 'Recipe not found' }, { status: 404 });
	}

	return json(recipe);
};
