<script lang="ts">
	export let orientation: 'horizontal' | 'vertical' = 'horizontal';
</script>

{#if orientation === 'horizontal'}
	<div class="bg-background-light-secondary dark:bg-background-dark-secondary h-px w-full"></div>
{:else}
	<div class="bg-background-light-secondary dark:bg-background-dark-secondary h-full w-px"></div>
{/if}
