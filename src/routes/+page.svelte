<script lang="ts">
	export let data;
</script>

<div>
	<textarea class="min-h-80 w-full border-none bg-background p-6">
		{JSON.stringify(data, null, 2)}
	</textarea>
</div>
