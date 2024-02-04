<script lang="ts">
	import { Avatar, Dropdown, DropdownDivider, DropdownItem } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';
	import { dev } from '$app/environment';

	const dispatch = createEventDispatcher();

	export let avatar_url: string | null = null;
	export let fallback_text: string | null = null;
</script>

<Avatar src={avatar_url ?? ''} class="hover:cursor-pointer">
	{fallback_text}
</Avatar>
<Dropdown>
	<a href="/user">
		<DropdownItem>Profile</DropdownItem>
	</a>
	<DropdownItem on:click={() => dispatch('signout')}>Log Out</DropdownItem>
	{#if dev}
		<DropdownDivider />
		<a href="/dev"><DropdownItem>Dev Menu</DropdownItem></a>
	{/if}
</Dropdown>
