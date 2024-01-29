<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';

	export let data;
</script>

<div>
	<DataDebug {data} />
</div>
