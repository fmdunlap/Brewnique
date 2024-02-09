<script lang="ts">
	import { submitComment } from '$lib/data/client/comment';
	import type { Comment } from '$lib/data/recipe';
	import CommentThread from './CommentThread.svelte';
	import NewCommentBar from './NewCommentBar.svelte';

	export let threads: Comment[] = [];
	export let maxDepth: number = 3;
	export let recipeId: string;

	async function handleCommentSubmit() {
		const commentResponse = await submitComment(recipeId, commentValue, null);
		if (commentResponse) threads.push(commentResponse);
		location.reload();
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
