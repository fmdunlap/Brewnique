<script lang="ts">
	import NavBar from '$lib/components/nav/NavBar.svelte';
	import { onMount } from 'svelte';
	import '../app.pcss';
	import { beforeNavigate, goto } from '$app/navigation';
	import { ModeWatcher } from 'mode-watcher';
	import type { LayoutData } from './$types';
	import * as Dialog from '$lib/components/ui/dialog';
	import { page } from '$app/stores';

	export let data: LayoutData;

	const { session } = data;

	function shouldRedirectToOnboarding() {
		return (
			session &&
			session.user &&
			session.user.onboardingStatus != 'COMPLETE' &&
			!$page.url.pathname.startsWith('/onboarding')
		);
	}

	beforeNavigate(async ({ cancel, to }) => {
		if (to?.route.id != '/onboarding' && (await shouldRedirectToOnboarding())) {
			console.log('redirecting to onboarding');
			cancel();
			await goto('/onboarding');
		}
	});

	onMount(async () => {
		if (await shouldRedirectToOnboarding()) {
			console.log('redirecting to onboardingm2');

			await goto('/onboarding');
		}
	});
</script>

<ModeWatcher />
<div class="flex min-h-screen flex-col">
	<NavBar
		loggedIn={session != null}
		avatarUrl={session == null
			? null
			: `https://cdn.brewnique.io/avatars/${session.user.userId}.svg`}
		fallbackText={session == null ? null : session.user.email.slice(0, 1)}
	/>
	<div class="mx-auto flex grow flex-col p-6 md:w-5/6 md:p-0 md:pb-2">
		<slot />
	</div>
	<div class="w-full bg-secondary">
		<p class="m-auto py-12 text-center">This is the footer</p>
	</div>
</div>

<Dialog.Root></Dialog.Root>
