<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import Separator from '$lib/components/ui/separator/separator.svelte';

	import SignInForm from './SignInForm.svelte';
	import SsoButtons from './SSOButtons.svelte';
	import type { SupabaseClient } from '@supabase/supabase-js';

	export let supabase: SupabaseClient;
</script>

<Dialog.Root>
	<Dialog.Trigger>Login
	</Dialog.Trigger>
	<Dialog.Content
		class="rounded-xl max-w-screen bg-amber-700 p-0 text-white md:w-2/3"
	>
	<div class="flex flex-col md:flex-row md:h-96">
		<div class="im-div m-auto flex flex-grow w-screen h-48 md:h-full md:w-1/2"/>
		<div class="m-auto w-2/3 flex flex-col gap-y-4 py-8 md:w-1/3 md:px-24">
			<SsoButtons {supabase} />
			<Separator />
			<SignInForm />
		</div>
	</div>
	</Dialog.Content>
</Dialog.Root>

<style>
	.im-div {
		background-image: url('$lib/assets/images/signinup-background.png');
		background-size: cover;
		background-position: center;
	}
</style>
