<script lang="ts">
	import RecipeCard from '$lib/components/RecipeCard.svelte';
	import SearchSidebar from '$lib/components/search-sidebar/SearchSidebar.svelte';
	import { Spinner } from 'flowbite-svelte';
	import type { SearchOptions } from './api/v1/recipes/filterOptions.js';
	import { onMount } from 'svelte';
	import { recipe } from '$src/schema.js';

	let recipes: (typeof recipe.$inferSelect)[] = [];
	let selectedOptions: SearchOptions | undefined;

	let loading: boolean = true;

	function selectedOptionsAsQueryParams() {
		let searchParams = new URLSearchParams();
		if (selectedOptions) {
			searchParams.set('sortBy', selectedOptions.sortBy);
			searchParams.set('minAbv', selectedOptions.filter.minAbv.toString());
			searchParams.set('maxAbv', selectedOptions.filter.maxAbv.toString());
			searchParams.set('minBatchSize', selectedOptions.filter.minBatchSize.toString());
			searchParams.set('maxBatchSize', selectedOptions.filter.maxBatchSize.toString());
			searchParams.set('ratings', selectedOptions.filter.rating.join(','));
		}
		console.log(searchParams.toString());
		return searchParams.toString();
	}

	function fetchRecipes() {
		loading = true;
		fetch(`/api/v1/recipes?${selectedOptionsAsQueryParams()}`, {
			method: 'GET'
		})
			.then((res) => {
				if (res.ok) {
					return res.json();
				}
				throw new Error('Failed to fetch recipes');
			})
			.then((data) => {
				console.log(data);
				recipes = data;
			})
			.catch((err) => {
				console.error(err);
			})
			.finally(() => {
				loading = false;
			});
	}

	onMount(() => {
		fetchRecipes();
	});
</script>

<div class="flex grow flex-row">
	<SearchSidebar
		bind:selectedOptions
		on:click={async () => {
			fetchRecipes();
		}}
	/>
	{#if loading}
		<Spinner class="m-auto" size="20" />
	{:else}
		<div class="grid w-full grid-cols-4 p-4">
			{#each recipes as recipe}
				<div class="px-1 py-0">
					<RecipeCard
						id={recipe.id}
						title={recipe.name ?? ''}
						rating={recipe.rating ?? 0}
						saved={false}
						image={recipe.images ? recipe.images[0] : ''}
						batch_size={recipe.batchSize ?? 0}
						og={recipe.originalGravity ?? 1.0}
						fg={recipe.finalGravity ?? 1.0}
					/>
				</div>
			{/each}
		</div>
	{/if}
</div>
