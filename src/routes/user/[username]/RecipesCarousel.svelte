<script lang="ts">
	import RecipeCard from '$lib/components/RecipeCard.svelte';
	import type { recipe } from '$src/schema';

	import { ArrowLeft, ArrowRight } from 'lucide-svelte';

	export let recipes: (typeof recipe.$inferSelect)[] = [];
	export let loggedIn = false;

	$: firstCardIndex = 0;
	let lastCardIndex = recipes ? recipes.length - 3 : 0;
</script>

<div class="flex flex-row justify-between gap-x-4">
	{#if recipes.length > 3}
		<button
			on:click={() => {
				firstCardIndex = Math.max(0, firstCardIndex - 1);
			}}
		>
			<ArrowLeft
				class="my-auto h-10 w-10 rounded-full p-1 hover:bg-slate-200 dark:hover:bg-slate-800"
			/>
		</button>
	{/if}
	<div class="flex grow flex-row gap-x-4">
		{#each recipes as recipe, i}
			{#if i >= firstCardIndex && i < firstCardIndex + 3}
				<RecipeCard
					id={recipe.id}
					title={recipe.name ?? ''}
					rating={recipe.rating ?? 0}
					saved={false}
					image={recipe.images ? recipe.images[0] : ''}
					batch_size={recipe.batchSize ?? 0}
					og={recipe.originalGravity ?? 1.0}
					fg={recipe.finalGravity ?? 1.0}
					{loggedIn}
				/>
			{/if}
		{/each}
	</div>

	{#if recipes.length > 3}
		<button
			on:click={() => {
				console.log('what');
				firstCardIndex = Math.min(Math.max(0, lastCardIndex), firstCardIndex + 1);
			}}
		>
			<ArrowRight
				class="my-auto h-10 w-10 rounded-full p-1 hover:bg-slate-200 dark:hover:bg-slate-800"
			/>
		</button>
	{/if}
</div>
