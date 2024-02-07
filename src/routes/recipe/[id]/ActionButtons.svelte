<script lang="ts">
	import { Bookmark, GitFork, MoreHorizontal, Share } from 'lucide-svelte';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	export let showMore = false;
</script>

<div class="mb-auto hidden flex-row gap-x-2 pl-4 md:flex">
	<button
		on:click={() => {
			dispatch('bookmark');
		}}
	>
		<Bookmark />
	</button>
	<button
		on:click={() => {
			dispatch('fork');
		}}
	>
		<GitFork />
	</button>
	<button
		on:click={() => {
			dispatch('share');
		}}
	>
		<Share />
	</button>
	{#if showMore}
		<button>
			<MoreHorizontal />
		</button>
	{/if}
</div>
