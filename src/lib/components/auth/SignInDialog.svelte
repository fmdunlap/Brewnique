<script lang="ts">
	import * as Dialog from '$lib/components/ui/dialog';
	import { Button } from '$lib/components/ui/button';
	import Separator from '$lib/components/ui/separator/separator.svelte';

	import SignInCardBackground from '$lib/assets/images/signinup-background.png';
	import SignInForm from './SignInForm.svelte';
	import SsoButtons from './SSOButtons.svelte';
	import type { SupabaseClient } from '@supabase/supabase-js';

	export let supabase: SupabaseClient;
</script>

<Dialog.Root>
	<Dialog.Trigger>
		<Button>Login</Button>
	</Dialog.Trigger>
	<Dialog.Content
		class="flex h-fit max-w-max flex-row rounded-xl bg-amber-700 p-0 text-white md:w-1/2"
	>
		<img src={SignInCardBackground} alt="Sign in card background" class="w-1/3" />
		<div class="mx-auto flex flex-col gap-y-4 py-8">
			<SignInForm />
			<Separator />
			<SsoButtons {supabase} />
		</div>
	</Dialog.Content>
</Dialog.Root>

<style>
	img {
		object-fit: cover;
	}
</style>
