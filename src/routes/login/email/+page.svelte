<script>
	import { Button } from 'flowbite-svelte';
	import Separator from '$lib/components/Separator.svelte';
	import EmailLoginForm from './EmailLoginForm.svelte';

	export let data;
</script>

<div
	class="mx-auto my-auto flex w-full flex-col gap-y-4 rounded-lg border-2 border-slate-200 bg-background-light-secondary p-8 pt-4 dark:border-slate-800 dark:bg-background-dark-secondary md:w-1/2"
>
	<h1 class="pb-4 pt-2 text-center text-2xl font-bold">Login with Email</h1>
	<EmailLoginForm data={data.form} />
	<button class="py-2 hover:underline">Forgot Password</button>

	<Separator orientation={'horizontal'} />

	<h1 class="text-center text-xl font-bold">Don't have an account?</h1>
	<a href="/login/email/sign-up" class="mx-auto w-full"
		><Button class="w-full" color="alternative">Sign Up</Button></a
	>
</div>
