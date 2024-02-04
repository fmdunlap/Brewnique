<script lang="ts">
	import { Button, Label, Textarea } from 'flowbite-svelte';
	import { superForm } from 'sveltekit-superforms/client';
	import type { PageData } from './$types';

	export let data: PageData;

	let skip = false;
	const { form, errors, enhance } = superForm(data.form, {
		dataType: 'json'
	});
</script>

<form method="POST" use:enhance class="flex flex-col gap-y-4">
	<Label for="bio">Bio</Label>
	<Textarea
		id="bio"
		name="bio"
		bind:value={$form.bio}
		aria-invalid={$form.bio ? 'true' : undefined}
	/>
	<input type="hidden" name="skip" bind:value={skip} />
	{#if $errors.bio}
		<span class="text-red-500">{$errors.bio}</span>
	{/if}
	<Button type="submit">Submit</Button>
	<Button
		variant="secondary"
		type="submit"
		on:click={() => {
			skip = true;
		}}
	>
		Skip
	</Button>
</form>
