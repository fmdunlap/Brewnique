<script>
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { Spinner } from 'flowbite-svelte';

	export let data;
</script>

<Spinner class="m-auto w-full" size={'24'} />
<DataDebug {data} />
