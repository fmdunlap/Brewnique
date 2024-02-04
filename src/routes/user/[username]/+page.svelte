<script>
	import RecipeCard from '$lib/components/RecipeCard.svelte';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { ArrowLeft, ArrowRight } from 'lucide-svelte';
	import UserAvatar from './UserAvatar.svelte';
	import RecipesCarousel from './RecipesCarousel.svelte';

	export let data;
</script>

<div class="mx-auto flex w-5/6 flex-col gap-y-4 pt-4">
	<div class="mx-auto">
		<UserAvatar
			userAvatarUrl={data.user?.avatarUrl ?? ''}
			username={data.user?.username ?? ''}
			userBio={data.user?.bio ?? ''}
			isLoggedInUser={data.session?.user.userId == data.user?.id}
			on:click={() => {
				console.log('edit');
			}}
		/>
	</div>

	<h1 class="text-4xl font-bold">Recipes</h1>
	{#if data.recipes && data.recipes != undefined && data.recipes.length > 0}
		<RecipesCarousel recipes={data.recipes} />
	{/if}
</div>

<DataDebug {data} />
