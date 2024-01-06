<script lang="ts">
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import AvatarMenu from './AvatarMenu.svelte';
	import { goto, preloadData, pushState } from '$app/navigation';
	import LoginPage from '../../routes/auth/login/+page.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import type { Database } from '$lib/types/supabaseDB';
	import DarkModeToggle from './DarkModeToggle.svelte';
	import Button from './ui/button/button.svelte';
	import SearchBarButton from './SearchBarButton.svelte';
	import SearchBar from './SearchBar.svelte';
	import type { LayoutData } from '../../routes/$types';

	export let data: LayoutData;

	const { supabase, session } = data;

	$: loginState = {
		supabase: supabase,
		session: session
	};

	$: loginDialogOpen = false;

	async function onLoginPressed(e: MouseEvent & { currentTarget: HTMLAnchorElement }) {
		if (e.metaKey || e.ctrlKey) return;
		e.preventDefault();

		const { href } = e.currentTarget;

		const result = await preloadData(href);
		if (result.type === 'loaded' && result.status === 200) {
			pushState('', { loginOpen: true });
			loginDialogOpen = true;
		} else {
			goto(href);
		}
	}

	function closeDialog() {
		loginDialogOpen = false;
	}
</script>

<div
	class="sticky top-0 flex min-w-full flex-row justify-between border-b-2 border-black px-4 py-4 dark:border-white"
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
			{#if session != null}
				<AvatarMenu {data} />
			{:else}
				<a href="/auth/login" on:click={onLoginPressed}>
					<Button variant="secondary">Log In</Button>
				</a>
			{/if}
		</div>
	</div>
</div>

<Dialog.Root
	open={loginDialogOpen}
	onOpenChange={(open) => {
		if (!open) {
			loginDialogOpen = false;
		}
	}}
>
	<Dialog.Content class="md:w-2/3 md:max-w-full md:p-0">
		<LoginPage on:close={closeDialog} data={loginState} />
	</Dialog.Content>
</Dialog.Root>
