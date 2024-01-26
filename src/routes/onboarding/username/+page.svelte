<script lang="ts">
	import { superForm } from 'sveltekit-superforms/client';
	import type { PageData } from './$types';
	import { Button, Input, Label } from 'flowbite-svelte';
	import SuperDebug from 'sveltekit-superforms/client/SuperDebug.svelte';

	export let data: PageData;

	let formEl: HTMLFormElement;

	const { form, errors, enhance } = superForm(data.userForm);
</script>

<form
	class="flex flex-col gap-y-4"
	method="post"
	bind:this={formEl}
	use:enhance
	on:submit={() => console.log(JSON.stringify($form, null, 2))}
>
	<Label for="username">Username</Label>
	<Input type="text" name="username" bind:value={$form.username} />
	{#if $errors.username}<span class="text-red-500">{$errors.username}</span>{/if}
	<Button type="submit">Submit</Button>
</form>

<SuperDebug data={$form} />
