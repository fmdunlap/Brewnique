<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import type { SupabaseClient } from '@supabase/supabase-js';
	import SignInDialog from './auth/SignInDialog.svelte';

	export let signedIn: boolean = false;
	export let supabase: SupabaseClient;
</script>

<div class="sticky top-0 flex min-w-full flex-row justify-between bg-orange-500 px-4 py-4">
	<div class="my-auto flex flex-row gap-x-4">
		<a href="/">Home</a>
		<a href="/about">About</a>
	</div>
	<div class="my-auto">
		{#if signedIn}
			<Button
				on:click={async () => {
					await supabase.auth.signOut();
				}}>Logout</Button
			>
		{:else}
			<SignInDialog {supabase} />
		{/if}
	</div>
</div>
