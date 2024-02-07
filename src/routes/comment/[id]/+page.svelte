<script lang="ts">
	import CommentThread from '$lib/components/comment/CommentThread.svelte';
	import { ChevronLeft } from 'lucide-svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div
	class="mx-auto my-2 flex flex-col gap-y-4 rounded-xl bg-background-light-secondary p-4 dark:bg-background-dark-secondary md:w-4/5"
>
	{#if data.comment}
		<a class="flex flex-row hover:underline" href={`/recipe/${data.comment.data.recipeId}`}
			><ChevronLeft /> Go to Recipe</a
		>
		<CommentThread comment={data.comment} maxDepth={10} />
	{:else}
		<p>No comment found!</p>
	{/if}
</div>
