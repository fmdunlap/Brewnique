<script lang="ts">
	import { Checkbox } from 'flowbite-svelte';

	export let checked = false;
</script>

<div class="flex flex-row gap-x-4">
	<p class="line-clamp-1 text-lg">Saved Recipes</p>
	<Checkbox bind:checked />
</div>
