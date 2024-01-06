<script lang="ts">
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import * as Dropdown from '$lib/components/ui/dropdown-menu';
	import * as Avatar from '$lib/components/ui/avatar';
	import { invalidateAll } from '$app/navigation';
	import type { LayoutData } from '../../routes/$types';

	export let data: LayoutData;

	const signOut = async () => {
		await data.supabase.auth.signOut();
		invalidateAll();
	};
</script>

<div class="h-10 w-10 rounded-full">
	<Dropdown.Root>
		<Dropdown.Trigger>
			{#if data.session != null}
				<Avatar.Root>
					<Avatar.Image src={data.user_profile?.data.avatar_url} alt="Avatar" />
					<Avatar.Fallback>{data.session.user.email?.slice(0, 1)}</Avatar.Fallback>
				</Avatar.Root>
			{:else}
				<Avatar.Root>
					<Avatar.Image src="/assets/images/default-avatar.png" alt="Avatar" />
					<Avatar.Fallback>?</Avatar.Fallback>
				</Avatar.Root>
			{/if}
		</Dropdown.Trigger>
		<Dropdown.Content class="w-72">
			<Dropdown.Label>Profile</Dropdown.Label>
			<Dropdown.Group>
				<Dropdown.Item>Some Item</Dropdown.Item>
				<Dropdown.Item>Another Item</Dropdown.Item>
				<Dropdown.Item><a href="/dev">Dev Menu</a></Dropdown.Item>
				<Dropdown.Item on:click={signOut}>Log Out</Dropdown.Item>
			</Dropdown.Group>
		</Dropdown.Content>
	</Dropdown.Root>
</div>
