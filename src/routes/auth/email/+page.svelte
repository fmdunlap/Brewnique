<script>
	import * as Card from '$lib/components/ui/card';
	import EmailLoginForm from './EmailLoginForm.svelte';

	export let data;
</script>

<Card.Root>
	<Card.Header>
		<Card.Title>Login</Card.Title>
	</Card.Header>
	<Card.Content>
		<EmailLoginForm data={data.form} />
	</Card.Content>
</Card.Root>
