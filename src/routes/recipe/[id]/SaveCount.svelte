<script lang="ts">
	import { ThumbsUp } from 'lucide-svelte';

	export let count: number = 0;
</script>

<div class="flex h-fit flex-row gap-x-2 px-2">
	<p class="line my-auto align-baseline font-bold leading-none">{count}</p>
	<ThumbsUp />
</div>
