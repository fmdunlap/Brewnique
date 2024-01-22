import { v4 as uuidv4 } from 'uuid';
import { S3 } from './s3';
import { PutObjectCommand } from '@aws-sdk/client-s3';
import { db } from './db';
import { recipe, recipeIngredient } from '$src/schema';
import { and, eq } from 'drizzle-orm';

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

export async function getRecipes(
	limit: number,
	offset: number,
	unpublished: boolean = false,
	fromUserId: string | null = null
) {
	const andPredicates = [];

	if (!unpublished) {
		andPredicates.push(eq(recipe.published, true));
	}

	if (fromUserId != null) {
		andPredicates.push(eq(recipe.ownerId, fromUserId));
	}

	const resultRecipes = await db
		.select()
		.from(recipe)
		.limit(limit)
		.offset(offset)
		.where(and(...andPredicates));

	console.log(JSON.stringify(resultRecipes));

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
