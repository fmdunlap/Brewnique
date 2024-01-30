<script lang="ts">
	import { page } from '$app/stores';
	import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper } from 'flowbite-svelte';
	import { type SearchSidebarOptions, DefaultSearchSidebarOptions } from './types';
	import BatchSizeCheckboxes from './BatchSizeCheckboxes.svelte';
	import AbvCheckboxes from './AbvCheckboxes.svelte';
	import RatingCheckboxes from './RatingCheckboxes.svelte';
	import SortByButtons from './SortByButtons.svelte';
	import { SearchIcon } from 'lucide-svelte';
	$: activeUrl = $page.url.pathname;

	export let selectedOptions: SearchSidebarOptions = DefaultSearchSidebarOptions;
</script>

<Sidebar {activeUrl} class="h-[100%]] flex flex-col">
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
				<BatchSizeCheckboxes bind:selectedBatchSizes={selectedOptions.filter.size} />
				<p class="text-lg">ABV %</p>
				<AbvCheckboxes bind:selectedAbvFilters={selectedOptions.filter.abv} />
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
