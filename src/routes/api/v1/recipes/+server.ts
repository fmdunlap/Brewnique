import type { RequestHandler } from './$types';
import { isSortByValue, type FilterOptions } from './filterOptions';
import { getRecipes, getUserSavedRecipes } from '$lib/data/recipe';
import { error } from '@sveltejs/kit';

function extractSortBy(searchParams: URLSearchParams) {
	const sortBy = searchParams.get('sortBy');
	if (sortBy && isSortByValue(sortBy)) {
		return sortBy;
	}
	return 'Newest';
}

function parseMinAbv(searchParams: URLSearchParams) {
	const minAbv = parseFloat(searchParams.get('minAbv') ?? '0');
	if (isNaN(minAbv)) {
		throw error(400, 'Invalid minAbv filter value');
	}
	if (minAbv < 0) {
		throw error(400, 'Invalid minAbv filter value. Min abv is less than 0');
	}
	if (minAbv > 100) {
		throw error(400, 'Invalid minAbv filter value. Min abv is greater than 100');
	}
	return minAbv;
}

function parseMaxAbv(searchParams: URLSearchParams) {
	const maxAbv = parseFloat(searchParams.get('maxAbv') ?? '100');
	if (isNaN(maxAbv)) {
		throw error(400, 'Invalid maxAbv filter value');
	}
	if (maxAbv < 0) {
		throw error(400, 'Invalid maxAbv filter value. Max abv is less than 0');
	}
	if (maxAbv > 100) {
		throw error(400, 'Invalid maxAbv filter value. Max abv is greater than 100');
	}
	return maxAbv;
}

function parseMinBatchSize(searchParams: URLSearchParams) {
	const batchSize = parseInt(searchParams.get('minBatchSize') ?? '0');
	if (isNaN(batchSize)) {
		throw error(400, 'Invalid batchSize filter value');
	}
	if (batchSize < 0) {
		throw error(400, 'Invalid batchSize filter value. Batch size is less than 0');
	}
	return batchSize;
}

function parseMaxBatchSize(searchParams: URLSearchParams) {
	const batchSize = parseInt(searchParams.get('maxBatchSize') ?? '0');
	if (isNaN(batchSize)) {
		throw error(400, 'Invalid batchSize filter value');
	}
	if (batchSize < 0) {
		throw error(400, 'Invalid batchSize filter value. Batch size is less than 0');
	}
	return batchSize;
}

function parseRating(searchParams: URLSearchParams) {
	const rating = searchParams.get('ratings') ?? '';

	console.log(rating);

	if (rating === '') {
		return [];
	}

	const ratingValues = rating.split(',').map((value) => {
		const ratingValue = parseInt(value);
		if (isNaN(ratingValue)) {
			throw error(400, 'Invalid rating filter value');
		}
		if (ratingValue < 0) {
			throw error(400, 'Invalid rating filter value. Rating is less than 0');
		}
		if (ratingValue > 5) {
			throw error(400, 'Invalid rating filter value. Rating is greater than 5');
		}
		return ratingValue;
	});
	return ratingValues;
}

function parseOnlySaved(searchParams: URLSearchParams) {
	const onlySaved = searchParams.get('onlySaved') ?? 'false';

	if (onlySaved == 'true') {
		return true;
	}

	return false;
}

function extractFilter(searchParams: URLSearchParams) {
	const minAbv = parseMinAbv(searchParams);
	const maxAbv = parseMaxAbv(searchParams);
	const minBatchSize = parseMinBatchSize(searchParams);
	const maxBatchSize = parseMaxBatchSize(searchParams);
	const rating = parseRating(searchParams);
	const onlySaved = parseOnlySaved(searchParams);

	return {
		minAbv,
		maxAbv,
		minBatchSize,
		maxBatchSize,
		rating,
		onlySaved
	} as FilterOptions;
}

export const GET: RequestHandler = async ({ url, locals }) => {
	const sortBy = extractSortBy(url.searchParams);
	const filter = extractFilter(url.searchParams);

	let recipes = await getRecipes(30, 0, true, null, sortBy, filter);

	if (filter.onlySaved) {
		const session = await locals.auth.validate();
		if (!session) {
			return new Response('Attempted to get saved recipe without session', { status: 401 });
		}
		const userSavedRecipes = await getUserSavedRecipes(session.user.userId);
		recipes = recipes.filter((row) => {
			return userSavedRecipes.includes(row.id);
		});
	}

	return new Response(JSON.stringify(recipes), { status: 200 });
};
