<script lang="ts">
	import { page } from '$app/stores';
	import { Sidebar, SidebarGroup, SidebarWrapper } from 'flowbite-svelte';
	import { type SearchSidebarOptions, DefaultSearchSidebarOptions } from './types';
	import BatchSizeCheckboxes from './BatchSizeCheckboxes.svelte';
	import AbvCheckboxes from './AbvCheckboxes.svelte';
	import RatingCheckboxes from './RatingCheckboxes.svelte';
	import SortByButtons from './SortByButtons.svelte';
	$: activeUrl = $page.url.pathname;

	export let selectedOptions: SearchSidebarOptions = DefaultSearchSidebarOptions;
</script>

<Sidebar {activeUrl} class="h-[100%]] flex flex-col">
	<SidebarWrapper class="min-h-full grow px-4">
		<SidebarGroup>
			<p class="text-xl font-bold">Sort By</p>
			<SortByButtons bind:sortBy={selectedOptions.sortBy} />
		</SidebarGroup>
		<SidebarGroup border>
			<p class="text-xl font-bold">Filter</p>
			<div class="flex flex-col ps-4">
				<p class="text-lg">Rating</p>
				<RatingCheckboxes bind:selectedRatings={selectedOptions.filter.rating} />
				<p class="text-lg">Batch Size</p>
				<BatchSizeCheckboxes bind:selectedBatchSizes={selectedOptions.filter.size} />
				<p class="text-lg">ABV %</p>
				<AbvCheckboxes bind:selectedAbvFilters={selectedOptions.filter.abv} />
			</div>
		</SidebarGroup>
	</SidebarWrapper>
</Sidebar>
