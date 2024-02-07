<script lang="ts">
	import type { Comment } from '$lib/data/recipe';
	import CommentThread from './CommentThread.svelte';
	import NewCommentBar from './NewCommentBar.svelte';

	export let threads: Comment[] = [];
	export let maxDepth: number = 3;
	export let recipeId: string;

	async function handleCommentSubmit() {
		const commentResponse = await fetch('/api/v1/comment', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				recipeId: recipeId,
				content: commentValue
			})
		});

		if (commentResponse.ok) {
			commentValue = '';
			const values = await commentResponse.json();
			threads.push({
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
			});
		} else {
			console.error('Failed to submit comment');
		}
	}

	$: commentValue = '';
</script>

<div class="flex flex-col gap-y-4">
	<NewCommentBar submitString="Comment" bind:value={commentValue} on:submit={handleCommentSubmit} />
	{#each threads as thread}
		<div
			class="w-full rounded-xl bg-background-light-secondary p-4 dark:bg-background-dark-secondary"
		>
			<CommentThread comment={thread} {maxDepth} />
		</div>
	{/each}
</div>
