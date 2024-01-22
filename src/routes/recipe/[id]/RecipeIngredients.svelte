<script lang="ts">
	import type { recipeIngredient } from '$src/schema';

	export let ingredients: (typeof recipeIngredient.$inferSelect)[];
</script>

<div class="flex flex-col gap-y-2">
	<h1 class="text-xl font-bold">Ingredients</h1>
	<div class="mx-auto flex w-5/6 flex-col gap-y-1">
		{#each ingredients as ingredient}
			<p>
				{ingredient.quantity}
				{ingredient.unit ? `${ingredient.unit} of` : ``}
				{ingredient.name}
			</p>
		{/each}
	</div>
</div>
