<script lang="ts">
	import type { SuperValidated } from 'sveltekit-superforms';
	import type { LoginFormSchema } from '$lib/types/forms';
	import type { Database } from '$lib/types/supabaseDB';
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import ContinueWithGoogle from '$lib/components/auth/ContinueWithGoogle.svelte';
	import ContinueWithGithub from '$lib/components/auth/ContinueWithGithub.svelte';
	import ContinueWithEmail from '$lib/components/auth/ContinueWithEmail.svelte';
	import Separator from '$lib/components/ui/separator/separator.svelte';

	export let data: {
		session: Session | null;
		supabase: SupabaseClient<Database>;
		form: SuperValidated<typeof LoginFormSchema> | null;
	};

	const { supabase } = data;
</script>

<div
	class="my-auto flex flex-col rounded-xl border-2 border-orange-900 bg-orange-700 shadow-xl md:flex-row"
>
	<div
		class="im-div m-auto flex w-full flex-grow rounded-t-xl md:rounded-l-xl md:rounded-tr-none"
	/>
	<div class="m-auto flex w-2/3 flex-col gap-y-4 py-8 md:px-8">
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
		<Separator />
		<ContinueWithEmail />
		<!-- {#if data.form != null}
			<SignInForm data={data.form} />
			<SuperDebug data={data.form} />
		{/if} -->
	</div>
</div>

<style>
	.im-div {
		height: calc(50vh);
		background-image: url('$lib/assets/images/signinup-background.png');
		background-size: cover;
		background-position: center;
	}
</style>
