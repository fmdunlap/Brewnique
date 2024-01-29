<script lang="ts">
	import NavBar from '$lib/components/nav/NavBar.svelte';
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import type { LayoutData } from './$types';
	import * as Dialog from '$lib/components/ui/dialog';

	export let data: LayoutData;

	const { session } = data;
</script>

<ModeWatcher />
<div class="bg-background-light-primary dark:bg-background-dark-primary flex min-h-screen flex-col">
	<NavBar
		loggedIn={session != null}
		avatarUrl={session == null ? null : session.user.avatarUrl}
		fallbackText={session == null ? null : session.user.email.slice(0, 1)}
	/>
	<div class="mx-auto flex w-full grow flex-col p-6 md:w-5/6 md:p-0 md:pb-2">
		<slot />
	</div>
	<div class="w-full">
		<p class="m-auto py-12 text-center">This is the footer</p>
	</div>
</div>

<Dialog.Root></Dialog.Root>
