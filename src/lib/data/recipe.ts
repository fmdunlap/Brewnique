import { v4 as uuidv4 } from 'uuid';
import { S3 } from './s3';
import { PutObjectCommand } from '@aws-sdk/client-s3';
import { db } from './db';
import { recipe, recipeIngredient } from '$src/schema';
import { and, eq, desc, asc, or, lte, gte } from 'drizzle-orm';
import {
	DEFAULT_FILTER_OPTIONS,
	type FilterOptions,
	type SortByValue
} from '$src/routes/api/v1/recipes/filterOptions';

export async function getRecipe(id: string) {
	const recipeResult = await db.select().from(recipe).where(eq(recipe.id, id));

	if (recipeResult.length < 1) {
		return null;
	}

	return recipeResult[0];
}

export async function getRecipeWithIngredients(id: string) {
	const recipeResult = await db.select().from(recipe).where(eq(recipe.id, id));

	if (recipeResult.length < 1) {
		return null;
	}

	const recipeIngredients = await db
		.select()
		.from(recipeIngredient)
		.where(eq(recipeIngredient.recipeId, recipeResult[0].id));

	return { ...recipeResult[0], ingredients: recipeIngredients };
}

export async function getAllRecipesByUser(userId: string) {
	return db.select().from(recipe).where(eq(recipe.ownerId, userId));
}

export async function getPublishedRecipesByUser(userId: string) {
	return db
		.select()
		.from(recipe)
		.where(and(eq(recipe.ownerId, userId), eq(recipe.published, true)));
}

export async function getUnpublishedRecipesByUser(userId: string) {
	return db
		.select()
		.from(recipe)
		.where(and(eq(recipe.ownerId, userId), eq(recipe.published, false)));
}

function getSortByOperator(sortBy: SortByValue) {
	switch (sortBy) {
		case 'NameAsc':
			return asc(recipe.name);
		case 'NameDesc':
			return desc(recipe.name);
		case 'RatingAsc':
			return asc(recipe.rating);
		case 'RatingDesc':
			return desc(recipe.rating);
		case 'Newest':
			return desc(recipe.createdAt);
		case 'Oldest':
			return asc(recipe.createdAt);
		default:
			return desc(recipe.createdAt);
	}
}

function getAbvFilterOperators(minAbv: number, maxAbv: number) {
	const operators = [];
	if (minAbv >= 0) {
		operators.push(gte(recipe.abv, minAbv));
	}
	if (maxAbv <= 100) {
		operators.push(lte(recipe.abv, maxAbv));
	}
	return and(...operators);
}

function getBatchSizeFilterOperators(minBatchSize: number, maxBatchSize: number) {
	const operators = [];
	if (minBatchSize >= 0) {
		operators.push(gte(recipe.batchSize, minBatchSize));
	}
	if (maxBatchSize <= 100) {
		operators.push(lte(recipe.batchSize, maxBatchSize));
	}
	return and(...operators);
}

function getRatingFilterOperators(ratingFilterValues: number[]) {
	const operators = [];
	for (const value of ratingFilterValues) {
		operators.push(eq(recipe.rating, value));
	}
	return or(...operators);
}

function getFilterOperators(filterOptions: FilterOptions) {
	const operators = [];

	operators.push(getAbvFilterOperators(filterOptions.minAbv, filterOptions.maxAbv));

	operators.push(
		getBatchSizeFilterOperators(filterOptions.minBatchSize, filterOptions.maxBatchSize)
	);

	if (filterOptions.rating.length > 0) {
		operators.push(getRatingFilterOperators(filterOptions.rating));
	}

	return and(...operators);
}

export async function getRecipes(
	limit: number,
	offset: number,
	unpublished: boolean = false,
	fromUserId: string | null = null,
	sortBy: SortByValue | null = 'Newest',
	filter: FilterOptions = DEFAULT_FILTER_OPTIONS
) {
	const andPredicates = [];

	if (!unpublished) {
		andPredicates.push(eq(recipe.published, true));
	}

	if (fromUserId != null) {
		andPredicates.push(eq(recipe.ownerId, fromUserId));
	}

	if (!sortBy) {
		sortBy = 'Newest';
	}

	const filterOperators = getFilterOperators(filter);

	const resultRecipes = await db
		.select()
		.from(recipe)
		.limit(limit)
		.offset(offset)
		.where(and(...andPredicates, filterOperators))
		.orderBy(getSortByOperator(sortBy));

	return resultRecipes;
}

export async function uploadRecipePhoto(file: File, recipeId: string) {
	const filetype = file.type.split('/')[1];
	const key = `recipes/${recipeId}/${uuidv4()}.${filetype}`;
	console.log('uploading to ' + key);
	S3.send(
		new PutObjectCommand({
			Bucket: 'brewnique',
			Key: key,
			Body: Buffer.from(await file.arrayBuffer()),
			ContentType: file.type,
			ContentLength: file.size
		})
	);
	return 'https://cdn.brewnique.io/' + key;
}
