<script lang="ts">
	import ImageUpload from '$lib/components/ImageUpload.svelte';
	import { Button, Modal } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';
	let clickOutsideModal = false;
	const dispatch = createEventDispatcher();
	export let b64imgs: string[] = [];
	export function openModal() {
		clickOutsideModal = true;
	}
	export function closeModal() {
		clickOutsideModal = false;
	}
</script>

<Modal title="Upload New Avatar" bind:open={clickOutsideModal} outsideclose>
	{#if b64imgs.length == 0}
		<ImageUpload showPreview={false} bind:b64imgs multiple={false} />
	{/if}
	{#if b64imgs.length > 0}
		<div class="flex flex-col gap-y-4">
			<div class="flex flex-col gap-y-4">
				<img
					class="mx-auto h-36 w-36 rounded-full object-cover"
					src={b64imgs[0]}
					alt="User Avatar"
				/>
			</div>
			<Button
				class="mx-auto w-2/3 rounded-full p-2"
				on:click={() => {
					dispatch('submit');
				}}
				color="primary">Save</Button
			>
			<Button
				class="mx-auto w-2/3 rounded-full p-2"
				color="alternative"
				on:click={() => (b64imgs = [])}>Cancel</Button
			>
		</div>
	{/if}
</Modal>
