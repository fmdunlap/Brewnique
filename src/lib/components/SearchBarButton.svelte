<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Search } from 'lucide-svelte';
	import * as Popover from './ui/popover';
	import SearchBar from './SearchBar.svelte';

	let popoverOpen = false;
</script>

<Button
	on:click={() => {
		popoverOpen = !popoverOpen;
	}}
	variant="outline"
	size="icon"
>
	<Search class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all" />
	<span class="sr-only">Toggle theme</span>
</Button>

<Popover.Root bind:open={popoverOpen}>
	<Popover.Trigger />
	<Popover.Content strategy="fixed" class="-ml-2 h-auto w-11/12" fitViewport={true}>
		<SearchBar />
	</Popover.Content>
</Popover.Root>
