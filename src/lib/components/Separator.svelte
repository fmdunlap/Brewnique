<script lang="ts">
	export let orientation: 'horizontal' | 'vertical' = 'horizontal';
</script>

{#if orientation === 'horizontal'}
	<div class="h-px w-full bg-gray-300"></div>
{:else}
	<div class="h-full w-px bg-gray-300"></div>
{/if}
