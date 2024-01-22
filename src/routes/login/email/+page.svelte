<script>
	import * as Card from '$lib/components/ui/card';
	import EmailLoginForm from './EmailLoginForm.svelte';

	export let data;
</script>

<Card.Root class="mx-auto my-auto w-full pt-4 md:w-1/2">
	<Card.Content>
		<h1 class="pb-4 pt-2 text-center text-2xl font-bold">Login with Email</h1>
		<EmailLoginForm data={data.form} />
	</Card.Content>
</Card.Root>
