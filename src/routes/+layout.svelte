<script lang="ts">
	import NavBar from '$lib/components/nav/NavBar.svelte';
	import { onMount } from 'svelte';
	import '../app.pcss';
	import { beforeNavigate, goto, invalidate } from '$app/navigation';
	import { ModeWatcher } from 'mode-watcher';
	import type { LayoutData } from './$types';
	import * as Dialog from '$lib/components/ui/dialog';
	import { page } from '$app/stores';
	import { getUserProfile } from '$lib/data/profile';
	import { Switch } from '$lib/components/ui/switch';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';

	export let data: LayoutData;

	const { supabase, session } = data;

	async function shouldRedirectToOnboarding() {
		if (!session) return false;

		const user_profile = await getUserProfile(session, supabase);
		if (!user_profile) return false;
		if (user_profile.onboarding_state == 'completed') return false;
		if ($page.url.pathname.startsWith('/onboarding')) return false;

		return true;
	}

	beforeNavigate(async ({ cancel, to }) => {
		if (to?.route.id != '/onboarding' && (await shouldRedirectToOnboarding())) {
			cancel();
			await goto('/onboarding');
		}
	});

	onMount(async () => {
		if (await shouldRedirectToOnboarding()) {
			await goto('/onboarding');
		}
	});

	onMount(() => {
		const {
			data: { subscription }
		} = supabase.auth.onAuthStateChange((event, _session) => {
			if (_session?.expires_at !== session?.expires_at) {
				invalidate('supabase:auth');
			}
		});

		return () => subscription.unsubscribe();
	});

	$: showData = false;
</script>

<ModeWatcher />
<div class="flex min-h-screen flex-col">
	<NavBar {data} />
	<div class="mx-auto flex grow flex-col p-6 md:w-5/6 md:p-0 md:pb-2">
		<slot />
	</div>
	<div class="w-full bg-secondary">
		<p class="m-auto py-12 text-center">This is the footer</p>
	</div>
</div>

<Dialog.Root></Dialog.Root>
