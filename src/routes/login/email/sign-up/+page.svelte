<script lang="ts">
	import EmailSignupForm from './EmailSignupForm.svelte';

	export let data;
</script>

<div
	class="bg-background outline-secondary my-auto flex w-full flex-col gap-y-4 rounded-xl p-6 outline outline-1 md:mx-auto md:w-1/2"
>
	<h1 class="text-center text-2xl font-bold">Sign-Up</h1>
	<EmailSignupForm data={data.form} />
	<button
		on:click={() => {
			history.back();
		}}
		class="py-2 hover:underline"
		type="submit">Cancel</button
	>
</div>
