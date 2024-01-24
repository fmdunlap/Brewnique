<script>
	import RecipeCard from '$lib/components/RecipeCard.svelte';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { ArrowLeft, ArrowRight } from 'lucide-svelte';
	import UserAvatar from './UserAvatar.svelte';

	export let data;
	$: firstCardIndex = 0;
	let lastCardIndex = data.recipes ? data.recipes.length - 3 : 0;
</script>

<div class="mx-auto flex w-5/6 flex-col gap-y-4 pt-4">
	<div class="mx-auto w-1/6">
		<UserAvatar
			userAvatarUrl={data.user?.avatarUrl ?? ''}
			username={data.user?.username ?? ''}
			userBio={data.user?.bio ?? ''}
			isLoggedInUser={data.session?.user.userId == data.user?.id}
		/>
	</div>

	<h1 class="text-4xl font-bold">Recipes {firstCardIndex}</h1>
	{#if data.recipes && data.recipes != undefined && data.recipes.length > 0}
		<div class="flex flex-row gap-x-4">
			<button
				on:click={() => {
					firstCardIndex = Math.max(0, firstCardIndex - 1);
				}}
			>
				<ArrowLeft
					class="my-auto h-10 w-10 rounded-full p-1 hover:bg-slate-200 dark:hover:bg-slate-800"
				/>
			</button>
			{#each data.recipes as recipe, i}
				{#if i >= firstCardIndex && i < firstCardIndex + 3}
					<div class="w-1/3">
						<RecipeCard
							id={recipe.id}
							title={recipe.name ?? ''}
							rating={recipe.rating ?? 0}
							saved={false}
							image={recipe.images ? recipe.images[0] : ''}
							batch_size={recipe.batchSize ?? 0}
							batch_unit={recipe.batchUnit ?? 'gal'}
							og={recipe.originalGravity ?? 1.0}
							fg={recipe.finalGravity ?? 1.0}
						/>
					</div>
				{/if}
			{/each}
			<button
				on:click={() => {
					console.log('what');
					firstCardIndex = Math.min(Math.max(0, lastCardIndex), firstCardIndex + 1);
				}}
			>
				<ArrowRight
					class="my-auto h-10 w-10 rounded-full p-1 hover:bg-slate-200 dark:hover:bg-slate-800"
				/>
			</button>
		</div>
	{/if}
</div>

<DataDebug {data} />