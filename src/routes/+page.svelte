<script lang="ts">
	export let data;
</script>

<div>
	<textarea class="min-h-80 w-full bg-red-600 p-6">
		{JSON.stringify(data, null, 2)}
	</textarea>
	<textarea class="min-h-80 w-full bg-red-600 p-6">
		{JSON.stringify(data, null, 2)}
	</textarea>
</div>
