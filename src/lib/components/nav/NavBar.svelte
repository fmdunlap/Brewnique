<script lang="ts">
	import AvatarMenu from './AvatarMenu.svelte';
	import { goto, preloadData, pushState } from '$app/navigation';
	import LoginPage from '../../../routes/login/+page.svelte';
	import * as Dialog from '$lib/components/ui/dialog';
	import DarkModeToggle from './DarkModeToggle.svelte';
	import { Button } from 'flowbite-svelte';
	import SearchBarButton from './SearchBarButton.svelte';
	import SearchBar from './SearchBar.svelte';
	import { page } from '$app/stores';

	export let avatarUrl: string | undefined | null = undefined;
	export let fallbackText: string | undefined | null = undefined;
	export let loggedIn: boolean;

	$: loginDialogOpen = $page.state.loginOpen;

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
	class="bg-background-light-primary dark:bg-background-dark-primary dark:border-background-dark-secondary sticky top-0 z-40 flex min-w-full flex-row justify-between border-b-2 px-4 py-4"
>
	<div class="my-auto flex flex-row gap-x-2">
		<Button variant="outline">LOGO</Button>
		<a href="/"><Button variant="link">Home</Button></a>
		{#if loggedIn}
			<Button
				variant="link"
				on:click={async () => {
					await goto('/recipe/new');
				}}>New Recipe</Button
			>
		{/if}
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
			{#if loggedIn}
				<AvatarMenu
					avatar_url={avatarUrl}
					fallback_text={fallbackText}
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
