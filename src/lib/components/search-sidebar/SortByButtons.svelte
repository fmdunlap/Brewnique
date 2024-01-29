<script lang="ts">
	import { SortByOptions, type SortByValue } from './types';

	export let sortBy: SortByValue;
</script>

<ul class="flex w-[100%] flex-col">
	{#each Object.values(SortByOptions) as option, i}
		<button
			class={'w-full rounded-lg py-1 ps-3 text-start hover:bg-slate-700 hover:underline' +
				(sortBy == option.value ? ' bg-slate-700 font-bold' : '')}
			on:click={() => {
				sortBy = option.value;
			}}
		>
			{option.label}
		</button>
	{/each}
</ul>
