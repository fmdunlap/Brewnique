<script lang="ts">
	import AvatarMenu from './AvatarMenu.svelte';
	import { goto, preloadData, pushState } from '$app/navigation';
	import LoginPage from '../../../routes/login/+page.svelte';
	import DarkModeToggle from './DarkModeToggle.svelte';
	import { Button } from 'flowbite-svelte';
	import SearchBarButton from './SearchBarButton.svelte';
	import SearchBar from './SearchBar.svelte';
	import { page } from '$app/stores';
	import Modal from '../Modal.svelte';

	export let avatarUrl: string | undefined | null = undefined;
	export let loggedIn: boolean;

	$: loginDialogOpen = $page.state.loginOpen;
	let modal: Modal;
	let goBack: boolean = true;

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
	class="sticky top-0 z-40 flex min-w-full flex-row justify-between border-b-2 bg-background-light-primary px-4 py-4 dark:border-background-dark-secondary dark:bg-background-dark-primary"
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
				<AvatarMenu avatar_url={avatarUrl} on:signout={async () => await goto('/login/logout')} />
			{:else}
				<a href="/login" on:click={onLoginPressed}>
					<Button variant="secondary">Log In</Button>
				</a>
			{/if}
		</div>
	</div>
</div>

<Modal
	bind:this={modal}
	bind:showModal={loginDialogOpen}
	on:close={() => {
		if (goBack) {
			history.back();
		}
	}}
	rounded="xl"
>
	<LoginPage
		showCloseIcon
		on:close={() => {
			modal.close();
		}}
		on:email={() => {
			goBack = false;
			modal.close();
			goBack = true;
		}}
	/>
</Modal>
