<script lang="ts">
	import { Checkbox, Input } from 'flowbite-svelte';

	export let minBatchSize = 0;
	export let maxBatchSize = 100;
</script>

<div class="grid grid-cols-2 gap-y-2 py-2 ps-2">
	<p class="my-auto">Min</p>
	<Input type="number" bind:value={minBatchSize} class="h-8" />
	<p class="my-auto">Max</p>
	<Input type="number" bind:value={maxBatchSize} class="h-8" />
</div>
