<script lang="ts">
	import type { SuperValidated } from 'sveltekit-superforms';
	import type { EmailSignupFormSchema } from './EmailSignupFormSchema';
	import { superForm } from 'sveltekit-superforms/client';
	import { Button, Input, Label } from 'flowbite-svelte';

	export let data: SuperValidated<EmailSignupFormSchema>;

	const { form, errors, constraints } = superForm(data);
</script>

<form method="POST" class="flex flex-col gap-y-4">
	<div class="flex flex-col gap-y-2">
		<Label for="email">Email</Label>
		<Input
			type="text"
			id="email"
			name="email"
			class="w-full"
			bind:value={$form.email}
			aria-invalid={$errors.email ? 'true' : undefined}
			{...$constraints.email}
		/>
		{#if $errors.email}
			{#each $errors.email as error}
				<span class="text-red-400">{error}</span>
			{/each}
		{/if}
	</div>
	<div class="flex flex-col gap-y-2">
		<Label for="password">Password</Label>
		<Input
			type="password"
			id="password"
			name="password"
			class="w-full"
			bind:value={$form.password}
			aria-invalid={$errors.password ? 'true' : undefined}
			{...$constraints.password}
		/>
		{#if $errors.password}
			{#each $errors.password as error}
				<span class="text-red-400">{error}</span>
			{/each}
		{/if}
	</div>
	<div class="flex flex-col gap-y-2">
		<Label for="confirmPassword">Confirm Password</Label>
		<Input
			type="password"
			id="confirmPassword"
			name="confirmPassword"
			class="w-full"
			bind:value={$form.confirmPassword}
			aria-invalid={$errors.confirmPassword ? 'true' : undefined}
			{...$constraints.confirmPassword}
		/>
		{#if $errors.confirmPassword}
			{#each $errors.confirmPassword as error}
				<span class="text-red-400">{error}</span>
			{/each}
		{/if}
	</div>
	<Button type="submit" class="mt-2">Sign Up</Button>
</form>
