<script lang="ts">
	import * as Card from '$lib/components/ui/card';
	import { Bookmark, Star, BookmarkCheck } from 'lucide-svelte';
	import { Separator } from '$lib/components/ui/separator';
	import { AspectRatio } from './ui/aspect-ratio';
	import BrewQuickFacts from './BrewQuickFacts.svelte';

	export let id: string;
	export let title: string;
	export let rating: number;
	export let saved: boolean;
	export let image: string;
	export let batch_size: number;
	export let batch_unit: string;
	export let og: number;
	export let fg: number;
</script>

<a href="/recipe/{id}">
	<Card.Root class="cursor-pointer transition-all hover:-translate-x-1 hover:-translate-y-1">
		<Card.Content>
			<div class="flex flex-col pt-4">
				<div class="flex max-h-min flex-row justify-between pb-2">
					<h1 class="line-clamp-1 overflow-clip text-xl font-bold">{title}</h1>
					<div class="my-auto flex flex-row gap-x-2">
						<Star class="dark:fill-white" />
						<p>{rating.toFixed(1)}</p>
						<button
							on:click={(e) => {
								e.preventDefault();
								saved = !saved;
							}}
						>
							{#if saved}
								<BookmarkCheck class="my-auto" />
							{:else}
								<Bookmark class="my-auto" />
							{/if}
						</button>
					</div>
				</div>
				<AspectRatio ratio={16/9}>
					<img class="object-cover rounded-md h-full w-full" src={image} alt="Recipe" />
				</AspectRatio>
				<BrewQuickFacts {batch_size} {batch_unit} {og} {fg} />
			</div>
		</Card.Content>
	</Card.Root>
</a>
