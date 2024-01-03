<script lang="ts">
	import Separator from '$lib/components/ui/separator/separator.svelte';

	import SignInForm from '$lib/components/auth/SignInForm.svelte';
	import SsoButtons from '$lib/components/auth/SSOButtons.svelte';
	import type { Session, SupabaseClient } from '@supabase/supabase-js';
	import type { Database } from '$lib/types/supabaseDB';
	import type { SuperValidated } from 'sveltekit-superforms';
	import type { LoginFormSchema } from '$lib/types/forms';

	export let data: {
		session: Session | null;
		supabase: SupabaseClient<Database>;
		form: SuperValidated<typeof LoginFormSchema> | null;
	};
</script>

<div
	class="my-auto flex flex-col rounded-xl border-2 border-orange-900 bg-orange-700 shadow-xl md:flex-row"
>
	<div
		class="im-div m-auto flex w-full flex-grow rounded-t-xl md:rounded-l-xl md:rounded-tr-none"
	/>
	<div class="m-auto flex w-2/3 flex-col gap-y-4 py-8 md:px-8">
		<SsoButtons supabase={data.supabase} />
		<Separator />
		{#if data.form}
			<SignInForm data={data.form} />
		{/if}
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
