<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import { Textarea } from '$lib/components/ui/textarea';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';
	import * as Card from '$lib/components/ui/card';
	import { superForm } from 'sveltekit-superforms/client';
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	export let data;

	const { user_profile } = data;

	const { form: displayNameForm, errors: displayNameErrors } = superForm(data.displayNameForm);
	const { form: bioForm, errors: bioFormErrors } = superForm(data.bioForm);

	let bioFormEl;

	let skip = false;

	onMount(() => {
		if (user_profile?.onboarding_state === 'completed') goto('/');
	});
</script>

<p class="mx-auto w-full py-10 text-center text-4xl">Welcome!</p>

{#if user_profile?.onboarding_state === 'email_unconfirmed'}
	<p>Check your email</p>
{/if}

<Card.Root>
	<Card.Content class="pt-4">
		{#if user_profile?.onboarding_state === 'display_name_pending'}
			<form class="flex flex-col gap-y-4" action="?/display_name" method="POST">
				<Label for="display_name">Display Name</Label>
				<Input
					type="text"
					id="display_name"
					name="display_name"
					bind:value={$displayNameForm.display_name}
					aria-invalid={$displayNameErrors.display_name ? 'true' : undefined}
				/>
				{#if $displayNameErrors.display_name}<span class="text-red-500"
						>{$displayNameErrors.display_name}</span
					>{/if}
				<Button type="submit">Submit</Button>
			</form>
		{/if}

		{#if user_profile?.onboarding_state === 'bio_pending'}
			<form method="POST" action="?/bio" bind:this={bioFormEl} class="flex flex-col gap-y-4">
				<Label for="bio">Bio</Label>
				<Textarea
					id="bio"
					name="bio"
					bind:value={$bioForm.bio}
					aria-invalid={$bioForm.bio ? 'true' : undefined}
				/>
				<input type="hidden" name="skip" bind:value={skip} />
				{#if $bioFormErrors.bio}<span class="text-red-500">{$bioFormErrors.bio}</span>{/if}
				<Button type="submit">Submit</Button>
				<Button
					variant="secondary"
					type="submit"
					on:click={() => {
						skip = true;
					}}>Skip</Button
				>
			</form>
		{/if}
	</Card.Content>
</Card.Root>
