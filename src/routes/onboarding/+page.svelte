<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import { superForm } from 'sveltekit-superforms/client';

	export let data;

	const { user_profile } = data;

	const { form, errors } = superForm(data.form);
</script>

<p class="mx-auto w-full py-10 text-center text-4xl">Welcome!</p>

{#if user_profile?.onboarding_state === 'email_unconfirmed'}
	<p>Check your email</p>
{/if}

{#if user_profile?.onboarding_state === 'display_name_pending'}
	<form class="flex flex-col gap-y-4" action="?/display_name" method="POST">
		<Label for="display_name">Display Name</Label>
		<Input
			type="text"
			id="display_name"
			name="display_name"
			bind:value={$form.display_name}
			aria-invalid={$errors.display_name ? 'true' : undefined}
		/>
		{#if $errors.display_name}<span class="text-red-500">{$errors.display_name}</span>{/if}
		<Button type="submit">Submit</Button>
	</form>
{/if}
