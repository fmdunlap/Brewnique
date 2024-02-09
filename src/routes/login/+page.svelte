<script lang="ts">
	import ContinueWithGoogle from '$lib/components/auth/ContinueWithGoogle.svelte';
	import ContinueWithGithub from '$lib/components/auth/ContinueWithGithub.svelte';
	import ContinueWithEmail from '$lib/components/auth/ContinueWithEmail.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import { createEventDispatcher } from 'svelte';
	import ContinueWithFacebook from '$lib/components/auth/ContinueWithFacebook.svelte';
	import { X } from 'lucide-svelte';

	const dispatch = createEventDispatcher();
	export let showCloseIcon: boolean = false;
</script>

<div class="my-auto flex flex-col rounded-xl shadow-xl md:flex-row">
	<div
		class="im-div m-auto flex w-full flex-grow rounded-t-xl md:rounded-l-xl md:rounded-tr-none"
	/>
	<div class="flex w-4/5 flex-col">
		{#if showCloseIcon}
			<button
				on:click={() => {
					dispatch('close');
				}}
			>
				<X class="m-3 ml-auto" />
			</button>
		{/if}
		<div class="m-auto flex flex-col gap-y-4 py-8 md:px-8">
			<ContinueWithGoogle
				on:click={async () => {
					await goto('/login/google');
				}}
			/>
			<ContinueWithGithub
				on:click={async () => {
					await goto('/login/github');
				}}
			/>
			<ContinueWithFacebook
				on:click={async () => {
					await goto('/login/facebook');
				}}
			/>
			<!-- Separator -->
			<div class="flex flex-row items-center gap-x-4">
				<div class="flex-grow border-t-2 border-gray-300" />
				<div class="text-gray-500">or</div>
				<div class="flex-grow border-t-2 border-gray-300" />
			</div>
			<ContinueWithEmail
				on:click={async () => {
					await goto('/login/email');
					await invalidateAll();
					dispatch('email');
				}}
			/>
		</div>
	</div>
</div>

<style>
	.im-div {
		height: calc(50vh);
		background-image: url('https://cdn.brewnique.io/signinup-background.png');
		background-size: cover;
		background-position: center;
	}
</style>
