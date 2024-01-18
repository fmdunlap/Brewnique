import { superValidate } from 'sveltekit-superforms/server';
import { NewRecipeFormSchema } from './NewRecipeForm';
import { fail } from '@sveltejs/kit';
import { db } from '$lib/data/db';
import { convertBase64ToFile } from '$lib/data/util';
import { uploadRecipePhoto } from '$lib/data/recipe';
import { v4 as uuidv4 } from 'uuid';
import { recipe } from '$src/schema';

export const load = async () => {
	// Server API:
	const form = await superValidate(NewRecipeFormSchema);

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
			return fail(401, { form });
		}

		console.log(session.user.userId);
		console.log(form.data.images[0].length);
		const recipeUUID = uuidv4();

		const imageUrls = [];

		console.log('batch', form.data.batchQuantity, form.data.batchUnit);

		for (let i = 0; i < form.data.images.length; i++) {
			const imageFile = convertBase64ToFile(form.data.images[i]);
			imageUrls.push(await uploadRecipePhoto(imageFile, recipeUUID));
		}

		await db.insert(recipe).values({
			id: recipeUUID.toString(),
			name: form.data.name,
			description: form.data.description,
			ownerId: session.user.userId,
			pictures: imageUrls,
			batchSize: form.data.batchQuantity,
			batchUnit: form.data.batchUnit
		});

		// Yep, return { form } here to
		return { form };
	}
};
