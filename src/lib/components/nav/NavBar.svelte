<script lang="ts">
	import AvatarMenu from './AvatarMenu.svelte';
	import { goto, preloadData, pushState } from '$app/navigation';
	import LoginPage from '../../../routes/login/+page.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import DarkModeToggle from './DarkModeToggle.svelte';
	import Button from '../ui/button/button.svelte';
	import SearchBarButton from './SearchBarButton.svelte';
	import SearchBar from './SearchBar.svelte';
	import type { LayoutData } from '../../../routes/$types';
	import { page } from '$app/stores';
	import { onMount } from 'svelte';

	export let data: LayoutData;

	$: session = data.session;
	$: loginDialogOpen = $page.state.loginOpen;
	let avatar_url: string | undefined | null;
	$: avatar_url;
	$: fallback_text = session?.user?.email?.slice(0, 1);

	onMount(async () => {
		if (!session) return;
		avatar_url = session?.user?.image;
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
</script>

<div
	class="sticky top-0 z-40 flex min-w-full flex-row justify-between border-b-2 bg-background px-4 py-4"
>
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
				<AvatarMenu
					{avatar_url}
					{fallback_text}
					on:signout={async () => await goto('/login/logout')}
				/>
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
		<LoginPage on:close={() => (loginDialogOpen = false)} />
	</Dialog.Content>
</Dialog.Root>
