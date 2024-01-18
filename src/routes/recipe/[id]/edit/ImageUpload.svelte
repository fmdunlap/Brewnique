<script lang="ts">
	import { Dropzone } from 'flowbite-svelte';
	import ImageUploadIcon from './ImageUploadIcon.svg';

	export let b64imgs: string[] = [];
	$: b64imgs;

	const handleChange = (event: Event) => {
		const eventTarget = event.target as HTMLInputElement;
		if (!eventTarget) return;
		const files = eventTarget.files;
		if (files && files.length > 0) {
			for (let i = 0; i < files.length; i++) {
				addFileItem(files[i]);
			}
		}
	};

	function removeFile(index: number) {
		b64imgs = b64imgs.filter((_, i) => i !== index);
	}

	function fileTypeIsImage(file: File) {
		return file.type == 'image/png' || file.type == 'image/jpeg' || file.type == 'image/webp';
	}

	function convertFileToBase64(file: File) {
		return new Promise<string>((resolve, reject) => {
			const reader = new FileReader();
			reader.readAsDataURL(file);
			reader.onload = (e) => {
				const converted = e.target?.result;
				if (!converted || typeof converted != 'string') {
					reject('Failed to convert file to base64');
					return;
				}
				resolve(converted);
			};
		});
	}

	function addDataTransferItem(dti: DataTransferItem) {
		if (dti.kind !== 'file') {
			return;
		}
		const file = dti.getAsFile();
		if (!file) {
			return;
		}
		addFileItem(file);
	}

	function addFileItem(file: File) {
		if (!fileTypeIsImage(file)) {
			return;
		}
		convertFileToBase64(file).then((b64) => {
			b64imgs = [...b64imgs, b64];
		});
	}

	const dropHandle = (event: DragEvent) => {
		event.preventDefault();

		if (!event.dataTransfer) return;

		if (event.dataTransfer.items) {
			[...event.dataTransfer.items].forEach((item) => {
				if (item.kind === 'file') {
					addDataTransferItem(item);
				}
			});
		}
	};
</script>

<div class="flex w-full grow flex-col">
	<div class="flex flex-row gap-x-6">
		<div class="my-auto flex w-1/4 flex-col">
			<h2 class="text-lg font-bold">Images</h2>
			<p>Add a thumbnail image. Whatever you think represents your brew best.</p>
		</div>
		<Dropzone
			on:drop={dropHandle}
			on:dragover={(event) => {
				event.preventDefault();
			}}
			on:change={(e) => handleChange(e)}
			multiple
		>
			<img src={ImageUploadIcon} alt="Upload Icon" class="mx-auto my-4 h-16 w-full" />
			<p class="text-xl">
				Drag and drop an image, or <span class="text-emerald-700 dark:text-emerald-400"
					>click to Browse</span
				>
			</p>
			<p>Supports PNG, JPG, and WEBP up to 5M.</p>
		</Dropzone>
	</div>

	<div class="grid grid-cols-4 gap-x-6">
		{#each b64imgs as image, i}
			<div class="mt-6 text-center">
				<img src={image} alt="Preview" class="h-full w-full rounded-xl object-cover" />
				<button
					class="text-red-700 dark:text-red-400"
					on:click={() => {
						removeFile(i);
					}}>Remove</button
				>
			</div>
		{/each}
	</div>
</div>
