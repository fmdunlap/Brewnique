<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import UserAvatar from './UserAvatar.svelte';
	import RecipesCarousel from './RecipesCarousel.svelte';
	import EditAvatarModal from './EditAvatarModal.svelte';

	export let data;

	let editAvatarModal: EditAvatarModal;
</script>

<div class="mx-auto flex w-5/6 flex-col gap-y-4 pt-4">
	<div class="mx-auto">
		<UserAvatar
			userAvatarUrl={data.user?.avatarUrl ?? ''}
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
	<EditAvatarModal bind:this={editAvatarModal} />
</div>

<DataDebug {data} />
