<script lang="ts">
	import RecipeCard from '$lib/components/RecipeCard.svelte';

	export let data;

	const { recipes } = data;

	$: showData = false;
</script>

<input type="checkbox" bind:checked={showData} />
{#if showData}
	<pre>
    <textarea class="h-96 w-full">
      {JSON.stringify(data, null, 2)}
    </textarea>
  </pre>
{/if}

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
				abv={(recipe.original_gravity - recipe.final_gravity) * 131.25}
				og={recipe.original_gravity}
				fg={recipe.final_gravity}
			/>
		</div>
	{/each}
</div>
