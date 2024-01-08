<script lang="ts">
	import { Switch } from '$lib/components/ui/switch';
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { dev } from '$app/environment';
	import { Bookmark, GitFork, MoreHorizontal, Share, ThumbsUp } from 'lucide-svelte';
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
	$: recipe_owner = data.recipe_owner;
	$: description = recipe.description ? recipe.description.split('\n') : '';
	$: notes = recipe.notes ? recipe.notes.split('\n') : '';


	$: pictures = recipe.pictures.map((picture: string) => {
		return {
			src: picture,
			alt: 'Brew Picture'
		};
	});
</script>

<div class="flex flex-row md:hidden justify-between pb-4">
		<SaveCount count={203} />
	<div class="flex flex-row pl-3">
		<Bookmark />
		<GitFork />
		<Share />
		<MoreHorizontal />
	</div>
</div>
<div class="flex flex-row w-full md:w-2/3 max-w-4xl mx-auto md:pt-8">
	<div class="hidden md:block">
		<SaveCount count={203} />
	</div>
	<div class="flex w-full flex-col gap-y-3">
		<BrewImagesCarousel images={pictures} />
		<h1 class="text-2xl font-bold">{data.recipe.name}</h1>
		<div class="flex flex-row w-full gap-x-1">
			<p>by</p>
			{#if recipe_owner.display_name}
			<UserLink display_name={recipe_owner.display_name} />
			{:else}
			<p class="line-through">deleted</p>
			{/if}
		</div>
		<div>
			{#each description as line}
				<p class="text-sm pb-1">{line}</p>
			{/each}
		</div>
		<div class="w-4/5 mx-auto py-4 md:w-2/3">
			<BrewQuickFacts og={recipe.original_gravity} fg={recipe.final_gravity} batch_size={recipe.batch_size} batch_unit={recipe.batch_unit} />
		</div>
		<RecipeIngredients ingredients={recipe.ingredients} />
		{#if recipe.process_steps}
			<RecipeProcess process={recipe.process_steps} />
		{/if}
		{#if recipe.notes}
			<div>
				<h1 class="text-xl font-bold pb-2">Notes</h1>
				{#each notes as line}
					<p class="pb-1">{line}</p>
				{/each}
			</div>
		{/if}
		<Separator orientation="horizontal" class="w-full mx-auto my-4" />
		<div>
			<h1 class="text-xl font-bold pb-2">Comments</h1>
		</div>
	</div>
	<div class="flex-row pl-3 hidden md:flex">
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
			<DataDebug label="recipe" data={data.recipe} />
			<DataDebug label="recipe owner" data={data.recipe_owner} />
		{/if}
	</div>
{/if}
