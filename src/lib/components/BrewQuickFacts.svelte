<script lang="ts">
	import { Separator } from '$lib/components/ui/separator';

	export let batch_size: number = 0;
	export let batch_unit: string = '';
	export let og: number = 0;
	export let fg: number = 0;

	$: abv = (og - fg) * 131.25;
</script>

<div class="flex h-20 flex-row justify-around">
	<div class="my-auto flex flex-col">
		<p class="mx-auto text-center text-sm">Batch Size</p>
		<p class="mx-auto">{batch_size} {batch_unit}</p>
	</div>
	<Separator orientation="vertical" />
	<div class="my-auto flex flex-col">
		<p class="mx-auto text-sm">ABV</p>
		<p>{abv.toFixed(1)}%</p>
	</div>
	<Separator orientation="vertical" />
	<div class="my-auto grid grid-flow-dense grid-cols-2">
		<p class=" w-min">OG</p>
		<p class="w-min font-bold">{og.toFixed(3)}</p>
		<p class="w-min">FG</p>
		<p class="w-min font-bold">{fg.toFixed(3)}</p>
	</div>
</div>
