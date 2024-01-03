<script lang="ts">
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import SignInDialog from './auth/SignInDialog.svelte';
	import AvatarMenu from './AvatarMenu.svelte';

	export let supabase: SupabaseClient;
	export let session: Session | null;
</script>

<div class="sticky top-0 flex min-w-full flex-row justify-between bg-orange-500 px-4 py-4">
	<div class="my-auto flex flex-row gap-x-4">
		<a href="/">Home</a>
		<a href="/about">About</a>
	</div>
	<div class="my-auto">
		{#if session != null}
			<AvatarMenu {supabase} {session} />
		{:else}
			<SignInDialog {supabase} />
		{/if}
	</div>
</div>
