<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import UserAvatar from './UserAvatar.svelte';
	import RecipesCarousel from './RecipesCarousel.svelte';
	import EditAvatarModal from './EditAvatarModal.svelte';
	import { goto, invalidate, invalidateAll } from '$app/navigation';
	import { error } from '@sveltejs/kit';

	export let data;

	let editAvatarModal: EditAvatarModal;
	let b64imgs: string[] = [];

	$: avatarUrl = `${data.user?.avatarUrl}?${Date.now()}` ?? '';
</script>

<div class="mx-auto flex w-5/6 flex-col gap-y-4 pt-4">
	<div class="mx-auto">
		<UserAvatar
			userAvatarUrl={avatarUrl}
			username={data.user?.username ?? ''}
			userBio={data.user?.bio ?? ''}
			isLoggedInUser={data.session?.user.userId == data.user?.id}
			on:click={() => {
				editAvatarModal.openModal();
			}}
		/>
	</div>

	<h1 class="text-4xl font-bold">Recipes</h1>
	{#if data.recipes && data.recipes != undefined && data.recipes.length > 0}
		<RecipesCarousel recipes={data.recipes} />
	{/if}
	<EditAvatarModal
		bind:this={editAvatarModal}
		bind:b64imgs
		on:submit={async () => {
			await fetch(`/api/v1/user/${data.session?.user.userId}/avatar`, {
				method: 'POST',
				headers: {
					'Content-Type': 'image/base64'
				},
				body: JSON.stringify({ b64img: b64imgs[0] })
			});
			editAvatarModal.closeModal();
			location.reload();
		}}
	/>
</div>

<DataDebug {data} />
