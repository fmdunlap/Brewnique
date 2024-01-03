<script lang="ts">
	import { Button } from '$lib/components/ui/button';

	export let data;
</script>

<h1 class="text-red-500">Welcome to SvelteKit</h1>
<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>

<textarea class="w-full">
	{JSON.stringify(data, null, 2)}
</textarea>
