<script lang="ts">
	import { Rating, Spinner, Star } from 'flowbite-svelte';
	import { onMount } from 'svelte';

	export let content: string = '';
	export let title: string = '';
	export let rating: number = 0;

	let user: null | {
		name: {
			first: string;
			last: string;
		};
		picture: {
			medium: string;
		};
	} = null;

	onMount(async () => {
		fetch('https://randomuser.me/api/', {
			method: 'GET',
			headers: {
				'Content-Type': 'application/json'
			}
		}).then((res) => res.json().then((data) => (user = data.results[0])));
	});
</script>

{#if user != null}
	<div class="flex flex-row gap-x-4 rounded-lg bg-slate-50 p-4 dark:bg-slate-800">
		<div class="flex h-fit w-1/3 flex-row gap-x-2">
			<img src={user.picture.medium} alt="User" class="my-auto h-8 w-8 rounded-full" />
			<p class="my-auto">{user.name.first}</p>
		</div>
		<div class="flex w-full grow flex-col">
			<div class="flex flex-row justify-between">
				<h2 class="my-auto text-xl font-bold">{title}</h2>
				<span class="flex flex-row rounded-lg bg-primary-500 px-2 py-1 font-bold text-white">
					<Star fillColor={'#FFF'} fillPercent={rating * 20.0} />
					{rating.toFixed(2)}
				</span>
			</div>
			<p class="py-2">{content}</p>
		</div>
	</div>
{:else}
	<Spinner />
{/if}
