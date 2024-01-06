<script lang="ts">
	import * as Dropdown from '$lib/components/ui/dropdown-menu';
	import * as Avatar from '$lib/components/ui/avatar';
	import { createEventDispatcher, onMount } from 'svelte';

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
			<Dropdown.Label>Profile</Dropdown.Label>
			<Dropdown.Group>
				<Dropdown.Item>Some Item</Dropdown.Item>
				<Dropdown.Item>Another Item</Dropdown.Item>
				<Dropdown.Item><a href="/dev">Dev Menu</a></Dropdown.Item>
				<Dropdown.Item on:click={() => dispatch('signout')}>Log Out</Dropdown.Item>
			</Dropdown.Group>
		</Dropdown.Content>
	</Dropdown.Root>
</div>
