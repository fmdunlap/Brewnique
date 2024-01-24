<script lang="ts">
	import { superForm } from 'sveltekit-superforms/client';
	import type { PageData } from './$types';
	import { Button, Input, Label } from 'flowbite-svelte';

	export let data: PageData;

	const { form, errors, enhance } = superForm(data.form);
</script>

<form class="flex flex-col gap-y-4" use:enhance method="POST">
	<Label for="username">Username</Label>
	<Input
		type="text"
		id="username"
		name="username"
		bind:value={$form.username}
		aria-invalid={$form.username ? 'true' : undefined}
	/>
	{#if $errors.username}<span class="text-red-500">{$errors.username}</span>{/if}
	<Button type="submit">Submit</Button>
</form>
