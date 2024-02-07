import { createRecipeComment, getComment, getRecipe } from '$lib/data/recipe';
import { getUserProfileById } from '$lib/data/user';
import type { RequestHandler } from '@sveltejs/kit';
import { z } from 'zod';

export const POST: RequestHandler = async ({ request, locals }) => {
	const session = await locals.auth.validate();

	if (!session) {
		return new Response('Unauthorized', { status: 401 });
	}

	const commentJson = await request.json();

	const schema = z.object({
		parentId: z.string().optional(),
		recipeId: z.string(),
		content: z.string()
	});

	let comment: {
		recipeId: string;
		content: string;
		parentId?: string | undefined;
	} | null = null;

	try {
		console.log(commentJson);
		comment = schema.parse(commentJson);
	} catch (e) {
		console.log(e);
		return new Response('Bad request', { status: 400 });
	}

	if (!(await getRecipe(comment.recipeId))) {
		return new Response('Recipe Not Found', { status: 404 });
	}
	if (comment.parentId && !(await getComment(comment.parentId))) {
		return new Response('Parent Comment Not Found', { status: 404 });
	}

	const commentValues = createRecipeComment(
		comment.recipeId,
		session.user.userId,
		comment.parentId ?? null,
		comment.content
	);

	return new Response(JSON.stringify({ ...commentValues, user: session.user }), { status: 201 });
};
