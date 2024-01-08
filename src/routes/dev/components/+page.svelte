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
				title={recipe.name}
				rating={recipe.rating}
				saved={false}
				image={recipe.pictures ? recipe.pictures[0] : 'http://placekitten.com/300/200'}
				batch_size={recipe.batch_size}
				batch_unit="gal"
				og={recipe.original_gravity}
				fg={recipe.final_gravity}
			/>
		</div>
	{/each}
</div>
