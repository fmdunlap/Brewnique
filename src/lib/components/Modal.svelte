<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { twMerge } from 'tailwind-merge';
	const dispatch = createEventDispatcher();

	export let showModal: boolean;

	let dialog: HTMLDialogElement;
	export let rounded: 'sm' | 'md' | 'lg' | 'xl' | 'none' = 'md';

	export const close = () => {
		dialog.close();
		dispatch('close');
	};

	$: if (dialog && showModal) dialog.showModal();
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
	bind:this={dialog}
	on:click|self={() => {
		close();
	}}
	class={twMerge('mx-auto w-full max-w-2xl p-0', rounded === 'none' ? '' : `rounded-${rounded}`)}
>
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div on:click|stopPropagation>
		<slot />
	</div>
</dialog>

<style>
	dialog[open] {
		animation: zoom 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
	}
	@keyframes zoom {
		from {
			transform: scale(0.95);
		}
		to {
			transform: scale(1);
		}
	}
	dialog[open]::backdrop {
		animation: fade 0.2s ease-out;
	}
	@keyframes fade {
		from {
			opacity: 0;
		}
		to {
			opacity: 1;
		}
	}
</style>
