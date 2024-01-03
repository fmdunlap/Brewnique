<script lang="ts">
	import { Button } from '$lib/components/ui/button';
	import type { SupabaseClient } from '@supabase/supabase-js';
	import ContinueWithGoogle from './ContinueWithGoogle.svelte';
	import { goto, invalidateAll } from '$app/navigation';
	import ContinueWithGithub from './ContinueWithGithub.svelte';

	export let supabase: SupabaseClient;
</script>

<div class="mx-auto flex w-full flex-col gap-y-4">
	<ContinueWithGoogle
		on:click={async () => {
			await supabase.auth.signInWithOAuth({ provider: 'google' });
		}}
	/>
	<ContinueWithGithub
		on:click={async () => {
			await supabase.auth.signInWithOAuth({ provider: 'github' });
		}}
	/>
</div>
