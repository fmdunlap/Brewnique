<script lang="ts">
	import { Toggle } from 'flowbite-svelte';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { dev } from '$app/environment';
	import { Bookmark, GitFork, MoreHorizontal, Share } from 'lucide-svelte';
	import SaveCount from './SaveCount.svelte';
	import BrewImagesCarousel from './BrewImagesCarousel.svelte';
	import UserLink from './UserLink.svelte';
	import BrewQuickFacts from '$lib/components/BrewQuickFacts.svelte';
	import RecipeIngredients from './RecipeIngredients.svelte';
	import RecipeProcess from './RecipeProcess.svelte';
	import Separator from '$lib/components/Separator.svelte';
	import ActionButtons from './ActionButtons.svelte';
	import CommentSection from '$lib/components/comment/CommentSection.svelte';

	export let data;

	$: showData = false;

	$: recipe = data.recipe;
	$: recipeOwner = data.recipeOwner;

	$: description = recipe.description ? recipe.description.split('\n') : '';
	$: notes = recipe.notes ? recipe.notes.split('\n') : '';

	$: images = recipe.images
		? recipe.images.map((picture: string) => {
				return {
					src: picture,
					alt: 'Brew Picture'
				};
			})
		: [];
</script>

<div class="flex flex-row justify-between pb-4 md:hidden">
	<SaveCount count={203} />
	<div class="flex flex-row">
		<Bookmark />
		<GitFork />
		<Share />
		{#if data.session?.user.userId == recipe.ownerId}
			<MoreHorizontal />
		{/if}
	</div>
</div>
<div class="mx-auto flex flex-row md:w-2/3 md:pt-8">
	<div class="hidden md:block">
		<SaveCount count={203} />
	</div>
	<div class="flex w-full flex-col gap-y-8">
		<BrewImagesCarousel {images} />
		<h1 class="text-2xl font-bold">{data.recipe.name}</h1>
		<div class="flex w-full flex-row gap-x-1">
			<p>by</p>
			{#if recipeOwner && recipeOwner.username}
				<UserLink display_name={recipeOwner.username} />
			{:else}
				<p class="line-through">deleted</p>
			{/if}
		</div>
		<div>
			{#each description as line}
				<p class="pb-1 text-sm">{line}</p>
			{/each}
		</div>
		<div class="mx-auto w-4/5 md:w-2/3">
			<BrewQuickFacts
				og={recipe.originalGravity ?? 1.0}
				fg={recipe.finalGravity ?? 1.0}
				batch_size={recipe.batchSize ?? 1}
			/>
		</div>
		<RecipeIngredients ingredients={recipe.ingredients} />
		{#if recipe.process}
			<RecipeProcess process={recipe.process} />
		{/if}
		{#if recipe.notes}
			<div class="flex flex-col gap-y-2">
				<h1 class="text-xl font-bold">Notes</h1>
				{#each notes as line}
					<p>{line}</p>
				{/each}
			</div>
		{/if}
		<Separator orientation="horizontal" />
		<h1 class="pb-2 text-xl font-bold">Comments</h1>
		<CommentSection threads={data.threads} recipeId={data.recipe.id} />
	</div>
	<ActionButtons showMore={data.session?.user.userId == recipe.ownerId} />
</div>

{#if dev}
	<div class="flex flex-col gap-y-4">
		<div class="flex flex-row gap-x-2">
			<p>Show Data:</p>
			<Toggle bind:checked={showData} />
		</div>

		{#if showData}
			<DataDebug label="recipe" {data} toggleable={false} />
		{/if}
	</div>
{/if}
