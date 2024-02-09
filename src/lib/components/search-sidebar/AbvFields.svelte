<script lang="ts">
	import { Input } from 'flowbite-svelte';

	export let minAbv = 0;
	export let maxAbv = 100;
</script>

<div class="grid grid-cols-2 gap-y-2 py-2 ps-2">
	<p class="my-auto">Min</p>
	<Input type="number" bind:value={minAbv} class="h-8" />
	<p class="my-auto">Max</p>
	<Input type="number" bind:value={maxAbv} class="h-8" />
</div>
