<script lang="ts">
	import { Button } from '$lib/components/ui/button';

	export let data;
</script>

<div>
	<textarea class="w-full min-h-80 p-6 bg-red-200\">
		{JSON.stringify(data, null, 2)}
	</textarea>
</div>
