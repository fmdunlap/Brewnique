<script lang="ts">
	import { Bookmark, BookmarkCheck } from 'lucide-svelte';
	import { AspectRatio } from './ui/aspect-ratio';
	import BrewQuickFacts from './BrewQuickFacts.svelte';
	import { Rating } from 'flowbite-svelte';

	export let id: string;
	export let title: string;
	export let rating: number;
	export let saved: boolean;
	export let image: string;
	export let batch_size: number;
	export let og: number;
	export let fg: number;
</script>

<a
	href="/recipe/{id}"
	class="h-fit w-full cursor-pointer rounded-xl border-2 bg-background-light-secondary p-4 transition-all hover:-translate-x-1 hover:-translate-y-1 hover:shadow-lg dark:border-slate-800 dark:bg-background-dark-secondary"
>
	<div class="flex flex-col gap-y-2 pt-4">
		<div class="flex max-h-min flex-row justify-between pb-2">
			<h1 class="line-clamp-1 overflow-clip text-xl font-bold">{title}</h1>
			<div class="my-auto flex flex-row gap-x-2">
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
		<AspectRatio ratio={16 / 9}>
			<img class="h-full w-full rounded-md object-cover" src={image} alt="Recipe" />
		</AspectRatio>
		<div class="flex flex-row pt-2">
			<Rating {rating} total={5} />
			<span class="mx-1.5 my-auto h-1 w-1 rounded-full bg-gray-500 dark:bg-gray-400" />
			<a
				href="/"
				class="my-auto text-sm font-medium text-gray-900 underline hover:no-underline dark:text-white"
			>
				73 reviews
			</a>
		</div>
		<BrewQuickFacts {batch_size} {og} {fg} />
	</div>
</a>
