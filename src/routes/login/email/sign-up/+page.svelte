<script lang="ts">
	import { Button, Card } from 'flowbite-svelte';
	import EmailSignupForm from './EmailSignupForm.svelte';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { navigating } from '$app/stores';

	export let data;
</script>

<div
	class="my-auto flex w-full flex-col gap-y-4 rounded-xl bg-background p-6 outline outline-1 outline-secondary md:mx-auto md:w-1/2"
>
	<h1 class="text-center text-2xl font-bold">Sign-Up</h1>
	<EmailSignupForm data={data.form} />
	<Button
		on:click={() => {
			history.back();
		}}
		type="submit"
		class="mt-2"
		color="alternative">Cancel</Button
	>
</div>
