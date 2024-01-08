<script lang="ts">
	import { Switch } from '$lib/components/ui/switch';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { dev } from '$app/environment';

	export let data;

	$: showData = false;
</script>

{#if dev}
	<div class="flex flex-col gap-y-4">
		<div class="flex flex-row gap-x-2">
			<p>Show Data:</p>
			<Switch bind:checked={showData} />
		</div>

		{#if showData}
			<DataDebug label="recipe" data={data.recipe} />
			<DataDebug label="recipe owner" data={data.recipe_owner} />
		{/if}
	</div>
{/if}
