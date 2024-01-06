<script lang="ts">
	import AvatarMenu from './AvatarMenu.svelte';
	import { goto, invalidateAll, preloadData, pushState } from '$app/navigation';
	import LoginPage from '../../routes/login/+page.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import DarkModeToggle from './DarkModeToggle.svelte';
	import Button from './ui/button/button.svelte';
	import SearchBarButton from './SearchBarButton.svelte';
	import SearchBar from './SearchBar.svelte';
	import type { LayoutData } from '../../routes/$types';
	import { page } from '$app/stores';
	import { getUserProfile } from '$lib/data/profile';
	import { onMount } from 'svelte';

	export let data: LayoutData;
	const { supabase } = data;

	$: session = data.session;
	$: loginDialogOpen = $page.state.loginOpen;
	let avatar_url: string | null;
	$: avatar_url;
	$: fallback_text = session?.user.email?.slice(0, 1);

	onMount(async () => {
		if (!session) return;
		const user_profile = await getUserProfile(session, supabase);
		if (!user_profile) return;
		avatar_url = user_profile.avatar_url;
		console.log(avatar_url);
	});

	async function onLoginPressed(e: MouseEvent & { currentTarget: HTMLAnchorElement }) {
		if (e.metaKey || e.ctrlKey) return;
		e.preventDefault();

		const { href } = e.currentTarget;

		const result = await preloadData(href);
		if (result.type === 'loaded' && result.status === 200) {
			pushState('/login', { loginOpen: true });
		} else {
			goto(href);
		}
	}

	async function signOut() {
		await data.supabase.auth.signOut();
		await invalidateAll();
		await goto('/');
	}
</script>

<div class="sticky top-0 flex min-w-full flex-row justify-between border-b-2 px-4 py-4">
	<div class="my-auto flex flex-row gap-x-2">
		<Button variant="outline">LOGO</Button>
		<a href="/"><Button variant="link">Home</Button></a>
	</div>
	<div class="hidden grow px-4 md:flex">
		<SearchBar />
	</div>
	<div class="my-auto flex flex-row gap-x-2">
		<div class="md:hidden">
			<SearchBarButton />
		</div>
		<DarkModeToggle />
		<div class="my-auto">
			{#if session}
				<AvatarMenu {avatar_url} {fallback_text} on:signout={signOut} />
			{:else}
				<a href="/login" on:click={onLoginPressed}>
					<Button variant="secondary">Log In</Button>
				</a>
			{/if}
		</div>
	</div>
</div>

<Dialog.Root
	open={loginDialogOpen}
	onOpenChange={(open) => {
		if (!open) history.back();
	}}
>
	<Dialog.Content class="md:w-2/3 md:max-w-full md:p-0">
		<LoginPage on:close={() => (loginDialogOpen = false)} {data} />
	</Dialog.Content>
</Dialog.Root>
