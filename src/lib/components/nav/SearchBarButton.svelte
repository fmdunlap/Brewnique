<script lang="ts">
	import { Button, Popover } from 'flowbite-svelte';
	import { Search } from 'lucide-svelte';
	import SearchBar from './SearchBar.svelte';

	let popoverOpen = false;
</script>

<Button
	on:click={() => {
		popoverOpen = !popoverOpen;
	}}
	variant="outline"
>
	<Search class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all" />
	<span class="sr-only">Toggle theme</span>
</Button>

<Popover>
	<SearchBar />
</Popover>
