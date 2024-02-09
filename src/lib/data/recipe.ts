import { v4 as uuidv4 } from 'uuid';
import { S3 } from './s3';
import { PutObjectCommand } from '@aws-sdk/client-s3';
import { db } from './db';
import { recipe, recipeComment, recipeIngredient, recipeSave, user } from '$src/schema';
import { and, eq, desc, asc, or, lte, gte } from 'drizzle-orm';
import {
	DEFAULT_FILTER_OPTIONS,
	type FilterOptions,
	type SortByValue
} from '$src/routes/api/v1/recipes/filterOptions';
import { getUserProfileById } from './user';

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

async function getRecipeComments(recipeId: string) {
	return await db
		.select()
		.from(recipeComment)
		.where(eq(recipeComment.recipeId, recipeId))
		.orderBy(asc(recipeComment.createdAt));
}

export async function getComment(commentId: string) {
	return await db.select().from(recipeComment).where(eq(recipeComment.id, commentId));
}

export type Comment = {
	parent: Comment | null;
	children: Comment[];
	user: typeof user.$inferSelect | null;
	data: typeof recipeComment.$inferSelect;
};

export async function getRecipeThreads(recipeId: string) {
	const allComments = await getRecipeComments(recipeId);
	const commentMap: Map<string, Comment> = new Map();
	allComments.forEach((comment) => {
		commentMap.set(comment.id, {
			parent: null,
			children: [],
			user: null,
			data: comment
		});
	});

	for (const comment of allComments) {
		const user = await getUserProfileById(comment.userId);
		const thisComment = commentMap.get(comment.id);
		if (!thisComment) {
			continue;
		}

		thisComment.user = user;

		commentMap.set(comment.id, thisComment);

		const parentId = comment.parentId;
		if (!parentId) {
			continue;
		}

		thisComment.parent = commentMap.get(parentId) ?? null;
		commentMap.get(parentId)?.children.push(thisComment);
	}

	const threads: Comment[] = [];

	commentMap.forEach((comment) => {
		if (comment.parent == null) {
			threads.push(comment);
		}
	});

	return threads;
}

export async function getThreadFromComment(commentId: string) {
	const commentMatches = await db
		.select()
		.from(recipeComment)
		.where(eq(recipeComment.id, commentId));
	if (commentMatches.length < 1) {
		return null;
	}

	const threads = await getRecipeThreads(commentMatches[0].recipeId);
	while (threads.length > 0) {
		const thread = threads.pop();
		if (thread?.data.id === commentId) {
			return thread;
		}
		threads.push(...(thread?.children ?? []));
	}
	return null;
}

export async function createRecipeComment(
	recipeId: string,
	userId: string,
	parentId: string | null,
	content: string
) {
	const commentValues = {
		id: uuidv4(),
		recipeId: recipeId,
		userId: userId,
		parentId: parentId,
		content: content
	};

	await db.insert(recipeComment).values(commentValues);

	return commentValues;
}

export async function getUserSavedRecipes(userId: string) {
	const userSavedRecipes = await db.select().from(recipeSave).where(eq(recipeSave.userId, userId));
	return userSavedRecipes.map((row) => {
		return row.recipeId;
	});
}

export async function recipeIsSavedByUser(userId: string, recipeId: string) {
	return (
		(
			await db
				.select()
				.from(recipeSave)
				.where(and(eq(recipeSave.userId, userId), eq(recipeSave.recipeId, recipeId)))
		).length > 0
	);
}

export async function saveRecipe(recipeId: string, userId: string) {
	if (!(await recipeIsSavedByUser(userId, recipeId)))
		await db.insert(recipeSave).values({
			id: uuidv4(),
			recipeId,
			userId
		});
}

export async function unsaveRecipe(recipeId: string, userId: string) {
	if (await recipeIsSavedByUser(userId, recipeId))
		await db
			.delete(recipeSave)
			.where(and(eq(recipeSave.recipeId, recipeId), eq(recipeSave.userId, userId)));
}
