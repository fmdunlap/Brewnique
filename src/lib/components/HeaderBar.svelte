<script lang="ts">
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import AvatarMenu from './AvatarMenu.svelte';
	import { goto, preloadData, pushState } from '$app/navigation';
	import LoginPage from '../../routes/login/+page.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import type { Database } from '$lib/types/supabaseDB';

	export let supabase: SupabaseClient<Database>;
	export let session: Session | null;

	$: loginState = {
		supabase: supabase,
		session: session,
		form: null
	};

	$: loginDialogOpen = loginState.form != null;

	async function onLoginPressed(e: MouseEvent & { currentTarget: HTMLAnchorElement }) {
		if (e.metaKey || e.ctrlKey) return;
		e.preventDefault();

		const { href } = e.currentTarget;

		const result = await preloadData(href);
		if (result.type === 'loaded' && result.status === 200) {
			loginState.form = result.data.form;
			pushState('', { loginOpen: true });
		} else {
			goto(href);
		}
	}
</script>

<div class="sticky top-0 flex min-w-full flex-row justify-between px-4 py-4 border-b-2 border-white">
	<div class="my-auto flex flex-row gap-x-4">
		<a href="/" class="text-white">Home</a>
		<a href="/about" class="text-white">About</a>
	</div>
	<div class="my-auto">
		{#if session != null}
			<AvatarMenu {supabase} {session} />
		{:else}
			<a href="/login" class="text-white" on:click={onLoginPressed}>Log In</a>
		{/if}
	</div>
</div>

<Dialog.Root
	open={loginDialogOpen}
	onOpenChange={(open) => {
		if (!open) {
			loginState.form = null;
			history.back();
		}
	}}
>
	<Dialog.Content class="md:w-2/3 md:max-w-full md:p-0">
		<LoginPage data={loginState} />
	</Dialog.Content>
</Dialog.Root>
