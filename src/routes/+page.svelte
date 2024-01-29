<script lang="ts">
	import SearchSidebar from '$lib/components/search-sidebar/SearchSidebar.svelte';
	import type { SearchSidebarOptions } from '$lib/components/search-sidebar/types';

	let selectedOptions: SearchSidebarOptions | undefined;
</script>

<div class="flex grow flex-row">
	<SearchSidebar bind:selectedOptions />
</div>
