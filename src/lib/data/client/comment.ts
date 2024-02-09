export async function submitComment(recipeId: string, content: string, parentId: string | null) {
	const commentBody = JSON.stringify(
		parentId
			? {
					recipeId: recipeId,
					parentId: parentId,
					content: content
				}
			: {
					recipeId: recipeId,
					content: content
				}
	);

	const commentResponse = await fetch('/api/v1/comment', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: commentBody
	});

	if (commentResponse.ok) {
		content = '';
		const values = await commentResponse.json();
		return {
			children: [],
			data: {
				userId: values.userId,
				parentId: values.parentId,
				content: values.content,
				createdAt: new Date(Date.now()),
				id: values.id,
				recipeId: values.recipeId,
				updatedAt: new Date(Date.now())
			},
			parent: null,
			user: values.user
		};
	} else {
		console.error('Failed to submit comment');
	}
}
