<script lang="ts">
	import ImageUploadIsland from './ImageUploadIsland.svelte';
	import ImagePreviews from './ImagePreviews.svelte';

	export let images: string[] = [];

	let testImages = [
		'https://placehold.co/300x400',
		'https://placehold.co/600x800',
		'https://placehold.co/500x1200',
		'https://placehold.co/600x400'
	];
</script>

<div class="flex w-full grow flex-col gap-y-6">
	<div class="flex flex-row gap-x-6">
		<div class="my-auto flex w-1/5 flex-col">
			<h2 class="text-lg font-bold">Images</h2>
			<p>Add a thumbnail image. Whatever you think represents your brew best.</p>
		</div>
		<div class="grow">
			<ImageUploadIsland on:browsefile={() => console.log('BROWSING')} />
		</div>
	</div>
	<ImagePreviews images={testImages} />
</div>
