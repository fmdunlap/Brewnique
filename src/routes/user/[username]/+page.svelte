<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import UserAvatar from './UserAvatar.svelte';
	import RecipesCarousel from './RecipesCarousel.svelte';
	import EditAvatarModal from './EditAvatarModal.svelte';
	import { goto, invalidate, invalidateAll } from '$app/navigation';
	import { error } from '@sveltejs/kit';
	import UserBio from './UserBio.svelte';
	import EditBioModal from './EditBioModal.svelte';

	export let data;

	let editAvatarModal: EditAvatarModal;
	let editBioModal: EditBioModal;
	let b64imgs: string[] = [];

	let avatarUrl = `${data.user?.avatarUrl}?${Date.now()}` ?? '';
	let userBio = data.user?.bio ?? '';
	let bioModalError = '';
</script>

<div class="mx-auto flex w-5/6 flex-col gap-y-4 pt-4">
	<div class="mx-auto">
		<UserAvatar
			userAvatarUrl={avatarUrl}
			username={data.user?.username ?? ''}
			isLoggedInUser={data.session?.user.userId == data.user?.id}
			on:click={() => {
				editAvatarModal.openModal();
			}}
		/>
	</div>
	<UserBio
		bio={data.user?.bio ?? ''}
		showEdit={data.session?.user.userId == data.user?.id}
		on:edit={() => {
			editBioModal.openModal();
		}}
	/>

	<h1 class="text-4xl font-bold">Recipes</h1>
	{#if data.recipes && data.recipes != undefined && data.recipes.length > 0}
		<RecipesCarousel recipes={data.recipes} />
	{/if}
</div>

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
		avatarUrl = `${avatarUrl}?${Date.now()}`;
		location.reload();
	}}
/>

<EditBioModal
	bind:this={editBioModal}
	bind:bio={userBio}
	bind:error={bioModalError}
	on:submit={async () => {
		const resp = await fetch(`/api/v1/user/${data.session?.user.userId}/bio`, {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ bio: userBio })
		});

		if (resp.status != 200) {
			bioModalError = await resp.text();
			return;
		}

		editBioModal.closeModal();
		bioModalError = '';
		if (data.user) {
			data.user.bio = userBio;
		}
	}}
/>

<DataDebug {data} />
