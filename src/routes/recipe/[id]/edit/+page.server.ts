import { superValidate } from 'sveltekit-superforms/server';
import { NewRecipeFormSchema } from './NewRecipeForm';
import { error, fail } from '@sveltejs/kit';
import { db } from '$lib/data/db';
import { convertBase64ToFile } from '$lib/data/util';
import { uploadRecipePhoto } from '$lib/data/recipe';
import { recipe, recipeIngredient } from '$src/schema';
import { eq } from 'drizzle-orm';
import { v4 as uuidv4 } from 'uuid';

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

export const actions = {
	default: async ({ request, locals }) => {
		const form = await superValidate(request, NewRecipeFormSchema);

		// Convenient validation check:
		if (!form.valid) {
			// Again, return { form } and things will just work.
			return fail(400, { form });
		}

		const session = await locals.auth.validate();

		if (!session || !session.user) {
			error(401, 'You must be logged in to edit a recipe');
		}

		const recipeEntry = await getRecipe(form.data.id);

		if (recipeEntry.ownerId != session.user.userId) {
			error(401, 'You do not own this recipe');
		}

		const imageUrls = [];

		if (!form.data.images) {
			form.data.images = [];
		}
		for (let i = 0; i < form.data.images.length; i++) {
			let imageUrl = '';
			if (form.data.images[i].startsWith('http')) {
				imageUrl = form.data.images[i];
			} else {
				const imageFile = convertBase64ToFile(form.data.images[i]);
				imageUrl = await uploadRecipePhoto(imageFile, recipeEntry.id);
			}
			imageUrls.push(imageUrl);
		}

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
				notes: form.data.notes
			})
			.where(eq(recipe.id, recipeEntry.id));

		const dbIngredients = await getRecipeIngredients(recipeEntry.id);

		for (let i = 0; i < dbIngredients.length; i++) {
			await db.delete(recipeIngredient).where(eq(recipeIngredient.id, dbIngredients[i].id));
		}

		for (let i = 0; i < form.data.ingredients.length; i++) {
			await db.insert(recipeIngredient).values({
				id: uuidv4(),
				recipeId: recipeEntry.id,
				name: form.data.ingredients[i].name,
				quantity: form.data.ingredients[i].quantity,
				unit: form.data.ingredients[i].unit
			});
		}

		// Yep, return { form } here to
		return { form };
	}
};