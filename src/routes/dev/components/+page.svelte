<script lang="ts">
	import RecipeCard from '$lib/components/RecipeCard.svelte';

	export let data;

	const { recipes } = data;
</script>

<div class="container grid grid-cols-1 md:grid-cols-3">
	{#each recipes as recipe}
		<div class="p-2">
			<RecipeCard
				id={recipe.id}
				title={recipe.name ?? 'Untitled'}
				rating={recipe.rating ?? 0}
				saved={false}
				image={recipe.images ? recipe.images[0] : 'http://placekitten.com/300/200'}
				batch_size={recipe.batchSize ?? 0}
				og={recipe.originalGravity ?? 1.0}
				fg={recipe.finalGravity ?? 1.0}
			/>
		</div>
	{/each}
</div>
