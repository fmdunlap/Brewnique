<script lang="ts">
	export let process: string[] = [];
</script>

<div class="flex flex-col gap-y-2">
	<h1 class="text-xl font-bold">Process</h1>
	<ol class="mx-auto flex w-5/6 list-decimal flex-col gap-y-1">
		{#each process as step}
			<li>{step}</li>
		{/each}
	</ol>
</div>
