<script lang="ts">
    export let process: string[] = [];

</script>

<h1 class="font-bold text-xl">Process</h1>
<ol class="flex flex-col gap-y-1 w-5/6 mx-auto list-decimal">
    {#each process as step}
        <li>{step}</li>
    {/each}
</ol>
