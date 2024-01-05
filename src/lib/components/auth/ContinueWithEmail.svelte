<script lang="ts">
	import { Mail } from 'lucide-svelte';
	import Button from '../ui/button/button.svelte';
</script>

<Button variant="secondary" on:click>
	<div class="my-auto flex w-full flex-row">
		<Mail class="mr-[12px] h-[20px] w-[20px] min-w-[20px]" />
		<span class="mx-auto">Continue with Email</span>
		<span class="sr-only">Continue with Email</span>
	</div>
</Button>
