<script lang="ts">
	import { Button, Modal, Textarea } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';
	let clickOutsideModal = false;
	const dispatch = createEventDispatcher();

	export function openModal() {
		clickOutsideModal = true;
	}
	export function closeModal() {
		clickOutsideModal = false;
	}

	export let bio: string;
	export let error: string = '';
</script>

<Modal title="Edit User Bio" bind:open={clickOutsideModal} outsideclose>
	<Textarea class="h-24 w-full" bind:value={bio} />
	{#if error}
		<p class="text-red-500">{error}</p>
	{/if}
	<div class="flex flex-row justify-end gap-x-2">
		<Button
			on:click={() => {
				dispatch('submit');
			}}
			color="primary"
		>
			Submit
		</Button>
		<Button
			color="alternative"
			on:click={() => {
				closeModal();
			}}
		>
			Cancel
		</Button>
	</div>
</Modal>
