<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';
	// @ts-ignore
	import autosize from 'svelte-autosize';

	const dispatch = createEventDispatcher();

	export let submitString = 'Comment';
	export let value: string = '';
</script>

<div class="relative flex flex-row gap-x-4">
	<textarea use:autosize rows={2} class="w-full rounded-xl p-2 text-black" bind:value />
	<Button
		class="mt-auto h-fit"
		on:click={() => {
			dispatch('submit');
		}}>{submitString}</Button
	>
</div>
