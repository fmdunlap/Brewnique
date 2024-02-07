<script lang="ts">
	import type { Comment } from '$lib/data/recipe';
	import UserComment from './UserComment.svelte';

	export let comment: Comment;
	export let maxDepth: number = 3;
</script>

<div class="flex flex-col gap-y-4">
	<div class="flex flex-col gap-y-2">
		<UserComment {comment} />
	</div>
	{#if comment.children && maxDepth > 1}
		<div class="flex flex-col gap-y-4 pl-10">
			{#each comment.children as child}
				<svelte:self comment={child} maxDepth={maxDepth - 1} />
			{/each}
		</div>
	{:else}
		<button class="w-fit text-sm text-gray-400 hover:cursor-pointer hover:underline"
			>See More...</button
		>
	{/if}
</div>
