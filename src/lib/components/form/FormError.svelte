<script lang="ts">
	export let errorMessages: string[] | null = [];
</script>

{#if errorMessages != null && errorMessages.length > 0}
	<div
		class="relative rounded border border-red-400 bg-red-100 px-4 py-3 text-red-700 dark:border-red-800 dark:bg-red-800 dark:text-red-100"
		role="alert"
	>
		<strong class="font-bold">Error</strong>
		<ul class="list-inside list-disc">
			{#each errorMessages as message}
				<li>{message}</li>
			{/each}
		</ul>
	</div>
{/if}
