<script lang="ts">
	import { type SuperValidated } from 'sveltekit-superforms';
	import type { EmailLoginFormSchema } from './EmailLoginFormSchema';
	import { superForm } from 'sveltekit-superforms/client';
	import { Button, Label, Input } from 'flowbite-svelte';

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
</form>
