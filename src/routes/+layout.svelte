<script lang="ts">
	import NavBar from '$lib/components/NavBar.svelte';
	import { onMount } from 'svelte';
	import '../app.pcss';
	import { beforeNavigate, goto, invalidate, onNavigate } from '$app/navigation';
	import { ModeWatcher } from 'mode-watcher';
	import type { LayoutData } from './$types';
	import { page } from '$app/stores';

	export let data: LayoutData;

	const { supabase, session, user_profile } = data;

	beforeNavigate((navigation) => {
		if (
			!session ||
			navigation.to?.route.id?.includes('auth/onboarding') ||
			user_profile?.data.onboarding_state == 'complete'
		)
			return;

		if (session && user_profile?.data.onboarding_state != 'complete') {
			navigation.cancel();
			goto('/auth/onboarding');
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
</script>

<ModeWatcher />
<div class="flex min-h-screen flex-col">
	<NavBar {data} />
	<div class="mx-auto flex grow flex-col p-10 md:w-4/5">
		<slot />
	</div>
</div>
