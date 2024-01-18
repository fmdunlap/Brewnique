import { db } from '$lib/data/db';
import { recipe } from '$src/schema';
import { fail, redirect } from '@sveltejs/kit';
import { v4 as uuidv4 } from 'uuid';

export const load = async ({ locals }) => {
	const newRecipeUUID = uuidv4();
	const session = await locals.auth.validate();

	if (!session || !session.user) {
		return fail(401);
	}

	const userId = session.user.userId;

	await db.insert(recipe).values({
		id: newRecipeUUID,
		ownerId: userId
	});

	redirect(302, `/recipe/${newRecipeUUID}/edit`);
};
