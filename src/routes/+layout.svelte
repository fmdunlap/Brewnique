<script>
	import HeaderBar from '$lib/components/HeaderBar.svelte';
	import { onMount } from 'svelte';
	import '../app.pcss';
	import { invalidate } from '$app/navigation';

	export let data;

	let { supabase, session } = data;
	$: ({ supabase, session } = data);

	onMount(() => {
		const {
			data: { subscription }
		} = supabase.auth.onAuthStateChange((event, _session) => {
			if (_session?.expires_at !== session?.expires_at) {
				invalidate('supabase:auth');
			}
		});

		return () => subscription.unsubscribe();
	});
</script>

<HeaderBar supabase={data.supabase} {session} />
<div class="mx-auto flex min-h-screen flex-col p-10 md:w-4/5">
	<slot />
</div>
