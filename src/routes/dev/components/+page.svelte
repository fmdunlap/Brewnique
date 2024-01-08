<script lang="ts">
	import RecipeCard from '$lib/components/RecipeCard.svelte';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { Switch } from '$lib/components/ui/switch';

	export let data;

	const { recipes } = data;

	$: showData = false;
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
