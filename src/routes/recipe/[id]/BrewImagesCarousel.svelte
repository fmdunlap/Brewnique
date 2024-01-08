<script lang="ts">
	import { Carousel } from 'flowbite-svelte';
	import type { HTMLImgAttributes } from 'svelte/elements';

	export let images: HTMLImgAttributes[] = [];
</script>

<div>
	<Carousel transition={null} {images} let:Controls let:Indicators>
		<Controls />
		<Indicators />
	</Carousel>
</div>
