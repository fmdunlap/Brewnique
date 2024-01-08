<script lang="ts">
	import type { Tables } from '$lib/types/supabaseDB';

	export let ingredients: Tables<'recipe_ingredient'>[];
</script>

<h1 class="text-xl font-bold">Ingredients</h1>
<div class="mx-auto flex w-5/6 flex-col gap-y-1">
	{#each ingredients as ingredient}
		<p>{ingredient.quantity} {ingredient.unit ? `${ingredient.unit} of` : ``} {ingredient.name}</p>
	{/each}
</div>
