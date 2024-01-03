<script lang="ts">
	import Separator from '$lib/components/ui/separator/separator.svelte';

	import SignInForm from '$lib/components/auth/SignInForm.svelte';
	import SsoButtons from '$lib/components/auth/SSOButtons.svelte';

	export let data;
</script>

<div
	class="my-auto flex flex-col rounded-xl border-2 border-orange-900 bg-orange-700 shadow-xl md:flex-row"
>
	<div class="im-div m-auto flex w-screen flex-grow rounded-l-xl" />
	<div class="m-auto flex w-2/3 flex-col gap-y-4 px-8 py-8">
		<SsoButtons supabase={data.supabase} />
		<Separator />
		<SignInForm data={data.form} />
	</div>
</div>

<style>
	.im-div {
		height: calc(50vh);
		width: 90%;
		background-image: url('$lib/assets/images/signinup-background.png');
		background-size: cover;
		background-position: center;
	}
</style>
