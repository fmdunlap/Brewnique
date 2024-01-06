<script lang="ts">
	import { type SuperValidated } from 'sveltekit-superforms';
	import type { EmailLoginFormSchema } from '$lib/types/forms';
	import { superForm } from 'sveltekit-superforms/client';
	import { Label } from '$lib/components/ui/label';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	export let data: SuperValidated<EmailLoginFormSchema>;

	const { form, errors, constraints } = superForm(data);
</script>

<form method="POST" class="flex flex-col gap-y-4">
	<Label for="email">Email</Label>
	<Input
		type="email"
		id="email"
		name="email"
		bind:value={$form.email}
		aria-invalid={$errors.email ? 'true' : undefined}
		{...$constraints.email}
	/>
	{#if $errors.email}<span class="invalid">{$errors.email}</span>{/if}
	<Label for="password">Password</Label>
	<Input
		type="password"
		id="password"
		name="password"
		bind:value={$form.password}
		aria-invalid={$errors.password ? 'true' : undefined}
		{...$constraints.password}
	/>
	{#if $errors.password}<span class="invalid">{$errors.password}</span>{/if}
	<Button type="submit">Sign In</Button>
	<Button variant="secondary">Sign Up</Button>
	<Button variant="link">Forgot Password</Button>
</form>
