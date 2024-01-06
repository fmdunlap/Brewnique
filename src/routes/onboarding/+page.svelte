<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import { Label } from '$lib/components/ui/label';

	export let data;

	const { supabase, session, user_profile } = data;
</script>

<p class="mx-auto w-full py-10 text-center text-4xl">Welcome!</p>

{JSON.stringify(session?.user.user_metadata, null, 2)}

{#if user_profile?.onboarding_state === 'email_unconfirmed'}
	<p>Check your email</p>
{/if}

{#if user_profile?.onboarding_state === 'display_name_pending'}
	<form class="flex flex-col gap-y-4">
		<Label>Let's get you a display name...</Label>
		<Input placeholder={'Display Name'} />
	</form>
{/if}
