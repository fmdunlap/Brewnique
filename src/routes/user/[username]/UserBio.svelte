<script lang="ts">
	import { Pencil } from 'lucide-svelte';
	import { createEventDispatcher } from 'svelte';
	const dispatch = createEventDispatcher();

	export let bio: string;
	export let showEdit: boolean;
</script>

<div class="mx-auto flex flex-row gap-x-2">
	<p class="my-auto text-center">{bio}</p>
	{#if showEdit}
		<button
			on:click={() => {
				dispatch('edit');
			}}
			class="rounded bg-gray-200 p-2 duration-200 ease-in-out hover:bg-gray-200 hover:duration-200 hover:ease-in-out dark:bg-gray-800 dark:hover:bg-gray-700"
		>
			<Pencil class="mx-auto h-4 w-4" />
		</button>
	{/if}
</div>
