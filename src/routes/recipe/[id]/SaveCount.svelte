<script lang="ts">
	import { ThumbsUp } from "lucide-svelte";

    export let count: number = 0;
</script>

<div class="flex flex-row h-fit gap-x-2 px-2">
    <p class="my-auto font-bold align-baseline line leading-none">{count}</p>
    <ThumbsUp/>
</div>