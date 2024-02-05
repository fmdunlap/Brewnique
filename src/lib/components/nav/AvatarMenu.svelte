<script lang="ts">
	import { Dropdown, DropdownDivider, DropdownItem } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';
	import { dev } from '$app/environment';
	import { mode, toggleMode } from 'mode-watcher';

	const dispatch = createEventDispatcher();

	export let avatar_url: string | null = null;
</script>

<img
	src={avatar_url + '?' + Date.now() ?? ''}
	alt="User Avatar"
	class="h-12 w-12 rounded-full object-cover transition duration-200 ease-in-out hover:cursor-pointer hover:saturate-150"
/>
<Dropdown>
	<a href="/user">
		<DropdownItem>Profile</DropdownItem>
	</a>
	{#if $mode == 'dark'}
		<DropdownItem on:click={toggleMode}>Light Mode</DropdownItem>
	{:else}
		<DropdownItem on:click={toggleMode}>Dark Mode</DropdownItem>
	{/if}
	<DropdownDivider />

	<DropdownItem on:click={() => dispatch('signout')}>Log Out</DropdownItem>
	{#if dev}
		<DropdownDivider />
		<a href="/dev"><DropdownItem>Dev Menu</DropdownItem></a>
	{/if}
</Dropdown>
