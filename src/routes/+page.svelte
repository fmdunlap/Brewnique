<script lang="ts">
	import RecipeCard from '$lib/components/RecipeCard.svelte';
	import SearchSidebar from '$lib/components/search-sidebar/SearchSidebar.svelte';
	import type { SearchSidebarOptions } from '$lib/components/search-sidebar/types';
	import { recipe } from '$src/schema';
	import { Spinner } from 'flowbite-svelte';

	export let data;

	let selectedOptions: SearchSidebarOptions | undefined;

	$: loading = data.recipes == undefined;
</script>

<div class="flex grow flex-row">
	<SearchSidebar
		bind:selectedOptions
		on:click={() => {
			fetch('/', {
				method: 'POST',
				body: JSON.stringify(selectedOptions)
			});
		}}
	/>
	{#if loading}
		<Spinner class="m-auto" size="20" />
	{:else}
		<div class="grid w-full grid-cols-4 p-4">
			{#each data.recipes as recipe}
				<div class="px-1 py-0">
					<RecipeCard
						id={recipe.id}
						title={recipe.name ?? ''}
						rating={recipe.rating ?? 0}
						saved={false}
						image={recipe.images ? recipe.images[0] : ''}
						batch_size={recipe.batchSize ?? 0}
						batch_unit={recipe.batchUnit ?? 'gal'}
						og={recipe.originalGravity ?? 1.0}
						fg={recipe.finalGravity ?? 1.0}
					/>
				</div>
			{/each}
		</div>
	{/if}
</div>
