<script lang="ts">
	import type { Tables } from "$lib/types/supabaseDB";

    export let ingredients: Tables<'recipe_ingredient'>[];
</script>

<h1 class="font-bold text-xl">Ingredients</h1>
<div class="flex flex-col gap-y-1 w-5/6 mx-auto">
    {#each ingredients as ingredient}
        <p>{ingredient.quantity} {ingredient.unit ? `${ingredient.unit} of` : `` } {ingredient.name}</p>
    {/each}
</div>