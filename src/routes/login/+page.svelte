<script lang="ts">
	import ContinueWithGoogle from '$lib/components/auth/ContinueWithGoogle.svelte';
	import ContinueWithGithub from '$lib/components/auth/ContinueWithGithub.svelte';
	import ContinueWithEmail from '$lib/components/auth/ContinueWithEmail.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import { createEventDispatcher } from 'svelte';
	import { page } from '$app/stores';
	import { signIn, signOut } from '@auth/sveltekit/client';

	const dispatch = createEventDispatcher();
</script>

<div class="my-auto flex flex-col rounded-xl border-2 shadow-xl md:flex-row">
	<div
		class="im-div m-auto flex w-full flex-grow rounded-t-xl md:rounded-l-xl md:rounded-tr-none"
	/>
	<div class="m-auto flex w-2/3 flex-col gap-y-4 py-8 md:px-8">
		<ContinueWithGoogle
			on:click={async () => {
				await signIn('google');
			}}
		/>
		<ContinueWithGithub
			on:click={async () => {
				await signIn('github');
			}}
		/>
		<Separator class="dark:bg-white" />
		<ContinueWithEmail
			on:click={async () => {
				await goto('/auth/email');
				await invalidateAll();
				dispatch('close');
			}}
		/>
	</div>
</div>

<style>
	.im-div {
		height: calc(50vh);
		background-image: url('$lib/assets/images/signinup-background.png');
		background-size: cover;
		background-position: center;
	}
</style>
