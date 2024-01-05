<script lang="ts">
	import type { SupabaseClient } from '@supabase/supabase-js';
	import ContinueWithGoogle from './ContinueWithGoogle.svelte';
	import ContinueWithGithub from './ContinueWithGithub.svelte';
	import ContinueWithEmail from './ContinueWithEmail.svelte';

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
	<ContinueWithEmail />
</div>
