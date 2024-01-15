<script lang="ts">
	import { AspectRatio } from '$lib/components/ui/aspect-ratio';

	export let images: string[] = [];
</script>

<div class="flex flex-row justify-between gap-x-6">
	{#each images as image}
		<div class="h-44 w-1/4 text-center">
			<img src={image} alt="Preview" class="h-full w-full rounded-xl object-cover" />
			<button class="text-red-700 dark:text-red-400">Remove</button>
		</div>
	{/each}
</div>
