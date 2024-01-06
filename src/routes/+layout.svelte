<script lang="ts">
	import NavBar from '$lib/components/NavBar.svelte';
	import { onMount } from 'svelte';
	import '../app.pcss';
	import { invalidate } from '$app/navigation';
	import { ModeWatcher } from 'mode-watcher';
	import type { LayoutData } from './$types';
	import * as Dialog from '$lib/components/ui/dialog';

	export let data: LayoutData;

	const { supabase, session } = data;

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

<Dialog.Root></Dialog.Root>
