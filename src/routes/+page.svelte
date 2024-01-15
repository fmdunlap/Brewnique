<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';

	export let data;

	// id: "",
	// ownerId: "",
	// createdAt: undefined,
	// updatedAt: undefined,
	// name: "",
	// description: null,
	// published: undefined,
	// difficulty: undefined,
	// brewType: null,
	// originalGravity: null,
	// finalGravity: null,
	// sweetenedGravity: null,
	// process: null,
	// rating: undefined,
	// batchSize: null,
	// batchUnit: null,
	// pictures: null,
	// notes: null
</script>

<div>
	<DataDebug {data} />
</div>
