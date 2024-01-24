<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import * as Card from '$lib/components/ui/card';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import EmailLoginForm from './EmailLoginForm.svelte';

	export let data;
</script>

<Card.Root class="mx-auto my-auto w-full pt-4 md:w-1/2">
	<Card.Content class="flex flex-col gap-y-4">
		<h1 class="pb-4 pt-2 text-center text-2xl font-bold">Login with Email</h1>
		<EmailLoginForm data={data.form} />
		<Button variant="link">Forgot Password</Button>

		<Separator orientation={'horizontal'} />

		<h1 class="text-center text-xl font-bold">Don't have an account?</h1>
		<a href="/login/email/sign-up" class="mx-auto w-full"
			><Button variant="secondary" class="w-full">Sign Up</Button></a
		>
	</Card.Content>
</Card.Root>
