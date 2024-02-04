<script lang="ts">
	import { SortByOptions, type SortByValue } from '$src/routes/api/v1/recipes/filterOptions';

	export let sortBy: SortByValue;
</script>

<div class="flex w-[100%] flex-col">
	{#each Object.values(SortByOptions) as option, i}
		<button
			class={'w-full rounded-lg py-1 ps-3 text-start hover:bg-slate-300 hover:underline dark:hover:bg-slate-700' +
				(sortBy == option.value ? ' bg-slate-300 font-bold dark:bg-slate-700' : '')}
			on:click={() => {
				sortBy = option.value;
			}}
		>
			{option.label}
		</button>
	{/each}
</div>
