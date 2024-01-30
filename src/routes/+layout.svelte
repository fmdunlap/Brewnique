<script lang="ts">
	import NavBar from '$lib/components/nav/NavBar.svelte';
	import '../app.pcss';
	import { ModeWatcher } from 'mode-watcher';
	import type { LayoutData } from './$types';
	import Separator from '$lib/components/Separator.svelte';

	export let data: LayoutData;

	const { session } = data;
</script>

<ModeWatcher />
<div class="flex min-h-screen flex-col bg-background-light-primary dark:bg-background-dark-primary">
	<NavBar
		loggedIn={session != null}
		avatarUrl={session == null ? null : session.user.avatarUrl}
		fallbackText={session == null ? null : session.user.email.slice(0, 1)}
	/>
	<div class="mx-auto flex w-full grow flex-col p-6 md:p-0">
		<slot />
	</div>
	<div class="w-full bg-background-light-secondary dark:bg-background-dark-secondary">
		<Separator />
		<p class=" m-auto py-6 text-center">This is the footer</p>
	</div>
</div>
