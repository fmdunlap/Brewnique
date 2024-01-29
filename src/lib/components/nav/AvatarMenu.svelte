<script lang="ts">
	import * as Dropdown from '$lib/components/ui/dropdown-menu';
	import { Avatar } from 'flowbite-svelte';
	import { createEventDispatcher } from 'svelte';
	import { dev } from '$app/environment';

	const dispatch = createEventDispatcher();

	export let avatar_url: string | null = null;
	export let fallback_text: string | null = null;
</script>

<div class="h-10 w-10 rounded-full">
	<Dropdown.Root>
		<Dropdown.Trigger>
			<Avatar src={avatar_url ?? ''}>
				{fallback_text}
			</Avatar>
		</Dropdown.Trigger>
		<Dropdown.Content class="w-72">
			<Dropdown.Group>
				<a href="/user">
					<Dropdown.Item>Profile</Dropdown.Item>
				</a>
				<Dropdown.Item on:click={() => dispatch('signout')}>Log Out</Dropdown.Item>
				{#if dev}
					<Dropdown.Separator />
					<a href="/dev"><Dropdown.Item>Dev Menu</Dropdown.Item></a>
				{/if}
			</Dropdown.Group>
		</Dropdown.Content>
	</Dropdown.Root>
</div>
