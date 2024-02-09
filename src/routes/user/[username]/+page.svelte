<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import UserAvatar from './UserAvatar.svelte';
	import RecipesCarousel from './RecipesCarousel.svelte';
	import EditAvatarModal from './EditAvatarModal.svelte';
	import UserBio from './UserBio.svelte';
	import EditBioModal from './EditBioModal.svelte';
	import RecipeCard from '$lib/components/RecipeCard.svelte';

	export let data;

	let editAvatarModal: EditAvatarModal;
	let editBioModal: EditBioModal;
	let b64imgs: string[] = [];

	let avatarUrl = `${data.user?.avatarUrl}?${Date.now()}` ?? '';
	let userBio = data.user?.bio ?? '';
	let bioModalError = '';
</script>

<div class="mx-auto flex w-full flex-col gap-y-4 pt-4 md:w-5/6">
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
		<div class="hidden lg:block">
			<RecipesCarousel recipes={data.recipes} loggedIn={data.session != null} />
		</div>
		<div class="flex flex-col gap-y-4 lg:hidden">
			{#each data.recipes as recipe}
				<RecipeCard
					id={recipe.id}
					title={recipe.name ?? ''}
					rating={recipe.rating ?? 0}
					saved={false}
					image={recipe.images ? recipe.images[0] : ''}
					batch_size={recipe.batchSize ?? 0}
					og={recipe.originalGravity ?? 1.0}
					fg={recipe.finalGravity ?? 1.0}
				/>
			{/each}
		</div>
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
