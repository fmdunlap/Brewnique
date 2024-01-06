<script lang="ts">
	import * as Dropdown from '$lib/components/ui/dropdown-menu';
	import * as Avatar from '$lib/components/ui/avatar';
	import { createEventDispatcher, onMount } from 'svelte';
	import { dev } from '$app/environment';

	const dispatch = createEventDispatcher();

	export let avatar_url: string | null = null;
	export let fallback_text: string | null = null;
</script>

<div class="h-10 w-10 rounded-full">
	<Dropdown.Root>
		<Dropdown.Trigger>
			<Avatar.Root>
				<Avatar.Image src={avatar_url} alt="Avatar" />
				<Avatar.Fallback>{fallback_text}</Avatar.Fallback>
			</Avatar.Root>
		</Dropdown.Trigger>
		<Dropdown.Content class="w-72">
			<Dropdown.Group>
				<Dropdown.Item>Profile</Dropdown.Item>
				<Dropdown.Item on:click={() => dispatch('signout')}>Log Out</Dropdown.Item>
				{#if dev}
					<Dropdown.Separator />
					<a href="/dev"><Dropdown.Item>Dev Menu</Dropdown.Item></a>
				{/if}
			</Dropdown.Group>
		</Dropdown.Content>
	</Dropdown.Root>
</div>
