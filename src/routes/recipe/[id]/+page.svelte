<script lang="ts">
	export let data;
</script>

<pre>
  <textarea class="h-96 w-full">
    {JSON.stringify(data.recipe, null, 2)}
  </textarea>
  <textarea class="h-96 w-full">
    {JSON.stringify(data.recipe_owner, null, 2)}
  </textarea>
</pre>
