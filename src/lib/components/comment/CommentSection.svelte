<script lang="ts">
	import type { Comment } from '$lib/data/recipe';
	import CommentThread from './CommentThread.svelte';

	export let threads: Comment[] = [];
	export let maxDepth: number = 3;

	console.log(threads);
</script>

<div class="flex flex-col gap-y-4">
	{#each threads as thread}
		<div
			class="w-full rounded-xl bg-background-light-secondary p-4 dark:bg-background-dark-secondary"
		>
			<CommentThread comment={thread} {maxDepth} />
		</div>
	{/each}
</div>
