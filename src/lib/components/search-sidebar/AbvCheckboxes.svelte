<script lang="ts">
	import { Checkbox } from 'flowbite-svelte';
	import { AbvFilterOptions, type AbvFilterValue } from './types';

	export let selectedAbvFilters: AbvFilterValue[] = [];
</script>

<div class="flex flex-col gap-y-2 ps-4">
	{#each AbvFilterOptions as abvFilterOption}
		<div class="flex flex-row">
			<Checkbox
				class="my-auto"
				checked={selectedAbvFilters.includes(abvFilterOption.value)}
				on:change={() => {
					if (selectedAbvFilters.includes(abvFilterOption.value)) {
						selectedAbvFilters = selectedAbvFilters.filter((v) => v != abvFilterOption.value);
					} else {
						selectedAbvFilters = [...selectedAbvFilters, abvFilterOption.value];
					}
				}}
			>
				<p class="my-auto">{abvFilterOption.label}</p>
			</Checkbox>
		</div>
	{/each}
</div>
