<script lang="ts">
	export let orientation: 'horizontal' | 'vertical' = 'horizontal';
</script>

{#if orientation === 'horizontal'}
	<div class="h-px w-full bg-background-light-secondary dark:bg-slate-800"></div>
{:else}
	<div class="h-full w-px bg-background-light-secondary dark:bg-slate-800"></div>
{/if}
