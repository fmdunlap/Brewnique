<script lang="ts">
	import { Checkbox } from 'flowbite-svelte';
	import { BatchSizeOptions, type BatchSizeValue } from './types';

	export let selectedBatchSizes: BatchSizeValue[] = [];
</script>

<div class="flex flex-col gap-y-2 ps-4">
	{#each BatchSizeOptions as batchSizeOption}
		<div class="flex flex-row">
			<Checkbox
				class="my-auto"
				checked={selectedBatchSizes.includes(batchSizeOption.value)}
				on:change={() => {
					if (selectedBatchSizes.includes(batchSizeOption.value)) {
						selectedBatchSizes = selectedBatchSizes.filter((v) => v != batchSizeOption.value);
					} else {
						selectedBatchSizes = [...selectedBatchSizes, batchSizeOption.value];
					}
				}}
			>
				<p class="my-auto">{batchSizeOption.label}</p>
			</Checkbox>
		</div>
	{/each}
</div>
