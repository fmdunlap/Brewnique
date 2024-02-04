<script lang="ts">
	import { page } from '$app/stores';
	import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper } from 'flowbite-svelte';
	import BatchSize from './BatchSizeFields.svelte';
	import AbvFields from './AbvFields.svelte';
	import RatingCheckboxes from './RatingCheckboxes.svelte';
	import SortByButtons from './SortByButtons.svelte';
	import { SearchIcon } from 'lucide-svelte';
	import {
		DEFAULT_SEARCH_OPTIONS,
		type SearchOptions
	} from '$src/routes/api/v1/recipes/filterOptions';
	import { createEventDispatcher } from 'svelte';

	export let selectedOptions: SearchOptions = DEFAULT_SEARCH_OPTIONS;

	const dispatch = createEventDispatcher();
	function dispatchChange() {
		dispatch('change', selectedOptions);
	}

	$: selectedOptions, dispatchChange();
</script>

<Sidebar class="h-[100%]] flex flex-col">
	<SidebarWrapper class="flex min-h-full grow flex-col px-4">
		<SidebarGroup>
			<p class="text-xl font-bold">Sort By</p>
			<SortByButtons bind:sortBy={selectedOptions.sortBy} />
		</SidebarGroup>
		<SidebarGroup border class="mb-auto">
			<p class="text-xl font-bold">Filter</p>
			<div class="flex flex-col ps-4">
				<p class="text-lg">Rating</p>
				<RatingCheckboxes bind:selectedRatings={selectedOptions.filter.rating} />
				<p class="text-lg">Batch Size</p>
				<BatchSize
					bind:minBatchSize={selectedOptions.filter.minBatchSize}
					bind:maxBatchSize={selectedOptions.filter.maxAbv}
				/>
				<p class="text-lg">ABV %</p>
				<AbvFields
					bind:minAbv={selectedOptions.filter.minAbv}
					bind:maxAbv={selectedOptions.filter.maxAbv}
				/>
			</div>
		</SidebarGroup>
		<SidebarGroup>
			<SidebarItem
				label="Apply"
				class="bg-primary-500 text-white hover:text-black dark:bg-primary-600"
				on:click
			>
				<svelte:fragment slot="icon">
					<SearchIcon class="h-5 w-5" />
				</svelte:fragment>
			</SidebarItem>
		</SidebarGroup>
	</SidebarWrapper>
</Sidebar>
