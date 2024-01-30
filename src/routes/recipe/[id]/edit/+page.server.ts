import { superValidate } from 'sveltekit-superforms/server';
import { NewRecipeFormSchema } from './NewRecipeForm';
import { error, fail, redirect } from '@sveltejs/kit';
import { db } from '$lib/data/db';
import { convertBase64ToFile } from '$lib/data/util';
import { uploadRecipePhoto } from '$lib/data/recipe';
import { recipe, recipeIngredient } from '$src/schema';
import { eq } from 'drizzle-orm';
import { v4 as uuidv4 } from 'uuid';
import type { Session } from 'lucia';
import type { SuperValidated } from 'sveltekit-superforms';

async function getRecipe(recipeId: string): Promise<typeof recipe.$inferSelect> {
	const recipeEntry = (await db.select().from(recipe).where(eq(recipe.id, recipeId)))[0];

	if (!recipeEntry) {
		throw error(404, 'Recipe not found');
	}

	return recipeEntry;
}

async function getRecipeIngredients(recipeId: string) {
	const recipeIngredients = await db
		.select()
		.from(recipeIngredient)
		.where(eq(recipeIngredient.recipeId, recipeId));

	return recipeIngredients;
}

export const load = async ({ params, locals }) => {
	const recipeId = params.id;
	const session = await locals.auth.validate();

	if (!session || !session.user) {
		error(401, 'You must be logged in to edit a recipe');
	}
	const recipeEntry = await getRecipe(recipeId);

	if (recipeEntry.ownerId != session.user.userId) {
		error(401, 'You do not own this recipe');
	}

	// Server API:
	const form = await superValidate(NewRecipeFormSchema);

	form.data.id = recipeId;
	form.data.name = recipeEntry.name ?? '';
	form.data.description = recipeEntry.description ?? '';
	form.data.images = recipeEntry.images ?? [];
	form.data.batchSize = recipeEntry.batchSize ?? 1;
	form.data.batchUnit = recipeEntry.batchUnit ?? 'gal';
	form.data.originalGravity = recipeEntry.originalGravity ?? 1.0;
	form.data.finalGravity = recipeEntry.finalGravity ?? 1.0;
	form.data.process = recipeEntry.process ?? [];
	form.data.notes = recipeEntry.notes ?? '';

	form.data.ingredients = (await getRecipeIngredients(recipeId)).map((ingredient) => {
		return {
			name: ingredient.name ?? '',
			quantity: ingredient.quantity ?? 0,
			unit: ingredient.unit ?? 'lb'
		};
	});

	// Unless you throw, always return { form } in load and form actions.
	return { form };
};

async function uploadImages(images: string[] | null, recipeId: string) {
	if (!images) {
		return [];
	}

	const imageUrls: string[] = [];

	for (let i = 0; i < images.length; i++) {
		let imageUrl = '';
		if (images[i].startsWith('http')) {
			imageUrl = images[i];
		} else {
			const imageFile = convertBase64ToFile(images[i]);
			imageUrl = await uploadRecipePhoto(imageFile, recipeId);
		}
		imageUrls.push(imageUrl);
	}
	return imageUrls;
}

async function updateRecipeIngredients(
	recipeId: string,
	ingredients: { name: string; quantity: number; unit: typeof recipeIngredient.$inferInsert.unit }[]
) {
	const dbIngredients = await getRecipeIngredients(recipeId);

	for (let i = 0; i < dbIngredients.length; i++) {
		await db.delete(recipeIngredient).where(eq(recipeIngredient.id, dbIngredients[i].id));
	}

	for (let i = 0; i < ingredients.length; i++) {
		await db.insert(recipeIngredient).values({
			id: uuidv4(),
			recipeId: recipeId,
			name: ingredients[i].name,
			quantity: ingredients[i].quantity,
			unit: ingredients[i].unit as typeof recipeIngredient.$inferSelect.unit
		});
	}
}

async function updateRecipe(
	form: SuperValidated<typeof NewRecipeFormSchema>,
	session: Session | null,
	published: boolean
): Promise<{ isError: boolean; errorCode: number } | null> {
	if (!session || !session.user) {
		throw error(401, 'You must be logged in to edit a recipe.');
	}

	// Convenient validation check:
	if (!form.valid) {
		return { isError: true, errorCode: 400 };
	}

	const recipeEntry = await getRecipe(form.data.id);
	if (recipeEntry.ownerId != session.user.userId) {
		throw error(401, 'You do not own this recipe. You can only edit recipes you own.');
	}

	const imageUrls = await uploadImages(form.data.images, recipeEntry.id);

	await db
		.update(recipe)
		.set({
			name: form.data.name,
			description: form.data.description,
			images: imageUrls,
			batchSize: form.data.batchSize,
			batchUnit: form.data.batchUnit,
			originalGravity: form.data.originalGravity,
			finalGravity: form.data.finalGravity,
			process: form.data.process.filter((step) => step.length > 0),
			notes: form.data.notes,
			published: published
		})
		.where(eq(recipe.id, recipeEntry.id));

	await updateRecipeIngredients(recipeEntry.id, form.data.ingredients);

	return null;
}

export const actions = {
	save: async ({ request, locals }) => {
		const form = await superValidate(request, NewRecipeFormSchema);
		const updateResult = await updateRecipe(form, await locals.auth.validate(), false);

		if (updateResult) {
			return fail(updateResult.errorCode, { form });
		}

		throw redirect(302, `/recipe/${form.data.id}`);
	},
	publish: async ({ request, locals }) => {
		const form = await superValidate(request, NewRecipeFormSchema);

		const updateResult = await updateRecipe(form, await locals.auth.validate(), true);

		if (updateResult) {
			return fail(updateResult.errorCode, { form });
		}
		// Yep, return { form } here to
		throw redirect(302, `/recipe/${form.data.id}`);
	}
};
