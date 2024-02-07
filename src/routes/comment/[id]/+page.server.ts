import { getThreadFromComment, type Comment } from '$lib/data/recipe.js';

export const load = async ({ params }) => {
	const id = params.id;
	const comment: Comment | null = await getThreadFromComment(id);

	return {
		comment: comment
	};
};
