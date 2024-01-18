<script lang="ts">
	import { Switch } from '$lib/components/ui/switch';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { dev } from '$app/environment';
	import { Bookmark, GitFork, MoreHorizontal, Share } from 'lucide-svelte';
	import SaveCount from './SaveCount.svelte';
	import BrewImagesCarousel from './BrewImagesCarousel.svelte';
	import UserLink from './UserLink.svelte';
	import BrewQuickFacts from '$lib/components/BrewQuickFacts.svelte';
	import RecipeIngredients from './RecipeIngredients.svelte';
	import RecipeProcess from './RecipeProcess.svelte';
	import { Separator } from '$lib/components/ui/separator';

	export let data;

	$: showData = false;

	$: recipe = data.recipe;
	$: recipe_owner = data.recipe_owner[0];

	console.log(JSON.stringify(recipe_owner, null, 2));

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
	<div class="flex flex-row pl-3">
		<Bookmark />
		<GitFork />
		<Share />
		<MoreHorizontal />
	</div>
</div>
<div class="mx-auto flex w-full max-w-4xl flex-row md:w-2/3 md:pt-8">
	<div class="hidden md:block">
		<SaveCount count={203} />
	</div>
	<div class="flex w-full flex-col gap-y-3">
		<BrewImagesCarousel {images} />
		<h1 class="text-2xl font-bold">{data.recipe.name}</h1>
		<div class="flex w-full flex-row gap-x-1">
			<p>by</p>
			{#if recipe_owner.username}
				<UserLink display_name={recipe_owner.username} />
			{:else}
				<p class="line-through">deleted</p>
			{/if}
		</div>
		<div>
			{#each description as line}
				<p class="pb-1 text-sm">{line}</p>
			{/each}
		</div>
		<div class="mx-auto w-4/5 py-4 md:w-2/3">
			<BrewQuickFacts
				og={recipe.originalGravity}
				fg={recipe.finalGravity}
				batch_size={recipe.batchSize}
				batch_unit={recipe.batchUnit}
			/>
		</div>
		<RecipeIngredients ingredients={recipe.ingredients} />
		{#if recipe.process_steps}
			<RecipeProcess process={recipe.process_steps} />
		{/if}
		{#if recipe.notes}
			<div>
				<h1 class="pb-2 text-xl font-bold">Notes</h1>
				{#each notes as line}
					<p class="pb-1">{line}</p>
				{/each}
			</div>
		{/if}
		<Separator orientation="horizontal" class="mx-auto my-4 w-full" />
		<div>
			<h1 class="pb-2 text-xl font-bold">Comments</h1>
		</div>
	</div>
	<div class="hidden flex-row pl-3 md:flex">
		<Bookmark />
		<GitFork />
		<Share />
		<MoreHorizontal />
	</div>
</div>

{#if dev}
	<div class="flex flex-col gap-y-4">
		<div class="flex flex-row gap-x-2">
			<p>Show Data:</p>
			<Switch bind:checked={showData} />
		</div>

		{#if showData}
			<DataDebug label="recipe" data={data.recipe} toggleable={false} />
			<DataDebug label="recipe owner" data={data.recipe_owner} toggleable={false} />
		{/if}
	</div>
{/if}
