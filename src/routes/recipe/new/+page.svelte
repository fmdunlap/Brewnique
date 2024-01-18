<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { Card } from '$lib/components/ui/card';
	import { Button } from '$lib/components/ui/button';
	import { Input, Textarea, Select } from 'flowbite-svelte';
	import { superForm } from 'sveltekit-superforms/client';
	import type { PageData } from './$types';

	import ImageUpload from './ImageUpload.svelte';
	import { onMount } from 'svelte';

	export let data: PageData;
	const { form, enhance } = superForm(data.form);

	const batchSizeUnits = ['gal', 'liter', 'barrel', 'cup', 'oz'];
	onMount(() => {
		$form.batchUnit = 'gal';
	});

	let unit: string = '';
</script>

<div class="mx-auto w-3/4">
	<h1 class="py-6 text-center text-2xl">Create New Recipe</h1>
	<Card class="my-2 p-8">
		<form
			class="flex flex-col gap-y-12"
			use:enhance
			method="post"
			on:submit={(formData) => {
				console.log(JSON.stringify(new FormData(formData.currentTarget)));
			}}
		>
			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="name">Name</label>
				<Input
					type="text"
					id="name"
					name="name"
					placeholder="The marvelous melon mead"
					bind:value={$form.name}
				/>
			</div>

			<ImageUpload bind:b64imgs={$form.images} />

			{#each $form.images as _, i}
				<input type="hidden" name={`images`} bind:value={$form.images[i]} />
			{/each}

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="description">Description</label>
				<Textarea id="description" name="description" bind:value={$form.description} />
			</div>

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="batchsize">Batch Size</label>
				<div class="flex flex-row gap-x-4">
					<Input
						class="w-1/3"
						type="number"
						id="batchQuantity"
						name="batchQuantity"
						placeholder="1"
						bind:value={$form.batchQuantity}
						on:change={() => {
							const formattedNumber = Number.parseInt($form.batchQuantity.toString());
							if (formattedNumber > 0) {
								$form.batchQuantity = formattedNumber;
							} else {
								$form.batchQuantity = 1;
							}
						}}
					/>

					<Select
						class="w-1/5"
						items={batchSizeUnits.map((unit) => {
							return { value: unit, name: unit };
						})}
						id="batchUnit"
						name="batchUnit"
						bind:value={$form.batchUnit}
					/>
				</div>
			</div>

			<Button
				type="submit"
				on:click={(e) => {
					console.log($form.images[0].length);
				}}>Save</Button
			>
			{unit}
		</form>
	</Card>
</div>

<DataDebug {data} />
