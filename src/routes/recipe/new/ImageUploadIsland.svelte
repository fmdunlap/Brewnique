<script lang="ts">
	import ImageUploadIcon from './ImageUploadIcon.svg';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();
</script>

<div
	class="flex grow flex-col rounded-xl p-6 text-center outline-dashed outline-1 outline-slate-300 dark:outline-slate-700"
>
	<img src={ImageUploadIcon} alt="Upload Icon" class="mx-auto" />
	<p class="text-xl">
		Drag and drop an image, or
		<button
			class="text-emerald-700 dark:text-emerald-400"
			on:click={() => {
				dispatch('browsefile');
			}}
		>
			Browse
		</button>
	</p>
	<p>Supports PNG, JPG, and WEBP up to 5M.</p>
</div>
