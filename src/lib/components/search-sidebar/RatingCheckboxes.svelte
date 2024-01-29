<script lang="ts">
	import { Checkbox } from 'flowbite-svelte';
	import { Star } from 'lucide-svelte';

	export let selectedRatings: number[] = [];
</script>

<div class="flex flex-col gap-y-2 ps-4">
	{#each [5, 4, 3, 2, 1] as rating}
		<div class="flex flex-row">
			<Checkbox
				class="my-auto"
				checked={selectedRatings.includes(rating)}
				on:change={() => {
					if (selectedRatings.includes(rating)) {
						selectedRatings = selectedRatings.filter((r) => r != rating);
					} else {
						selectedRatings = [...selectedRatings, rating];
					}
					selectedRatings.sort((a, b) => b - a);
				}}
			>
				{#each [1, 2, 3, 4, 5] as star}
					<Star class={'my-auto stroke-1' + (star <= rating ? ' fill-primary-500' : '')} />
				{/each}
			</Checkbox>
		</div>
	{/each}
</div>
