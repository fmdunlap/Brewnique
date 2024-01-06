<script>
	import * as Card from '$lib/components/ui/card';
	import EmailLoginForm from './EmailLoginForm.svelte';

	export let data;
</script>

<Card.Root class="mx-auto w-full pt-4 md:w-1/2">
	<Card.Content>
		<EmailLoginForm data={data.form} />
	</Card.Content>
</Card.Root>
