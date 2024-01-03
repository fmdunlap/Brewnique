<script lang="ts">
	export let data;
</script>

<div>
	<textarea class="bg-red-200\ min-h-80 w-full p-6">
		{JSON.stringify(data, null, 2)}
	</textarea>
</div>
