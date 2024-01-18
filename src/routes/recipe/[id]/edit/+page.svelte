<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { Card } from '$lib/components/ui/card';
	import { Button } from '$lib/components/ui/button';
	import { Input, Textarea, Select } from 'flowbite-svelte';
	import { superForm } from 'sveltekit-superforms/client';

	import ImageUpload from './ImageUpload.svelte';
	import { onMount } from 'svelte';
	import type { PageData } from './$types';
	import { X } from 'lucide-svelte';

	export let data: PageData;
	const { form, enhance } = superForm(data.form, {
		dataType: 'json'
	});

	const batchSizeUnits = ['gal', 'liter', 'barrel', 'cup', 'oz'];
</script>

<div class="mx-auto w-3/4">
	<h1 class="py-6 text-center text-2xl">Create New Recipe</h1>
	<Card class="my-2 p-8">
		<form class="flex flex-col gap-y-12" use:enhance method="post">
			<!-- Hidden ID -->

			<input type="hidden" bind:value={$form.id} />

			<!-- Name -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="name">Name</label>
				<Input type="text" placeholder="The marvelous melon mead" bind:value={$form.name} />
			</div>

			<!-- Images -->

			<ImageUpload bind:b64imgs={$form.images} />

			{#each $form.images as _, i}
				<input type="hidden" bind:value={$form.images[i]} />
			{/each}

			<!-- Description -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="description">Description</label>
				<Textarea id="description" bind:value={$form.description} />
			</div>

			<!-- Batch Size -->

			<div class="flex w-1/2 flex-col gap-y-4">
				<label class="text-lg" for="batchSize">Batch Size</label>
				<div class="flex flex-row gap-x-4">
					<Input
						type="number"
						placeholder="1"
						bind:value={$form.batchSize}
						on:change={() => {
							const formattedNumber = Number.parseInt($form.batchSize.toString());
							if (formattedNumber > 0) {
								$form.batchSize = formattedNumber;
							} else {
								$form.batchSize = 1;
							}
						}}
					/>

					<Select
						class="w-1/5"
						items={batchSizeUnits.map((unit) => {
							return { value: unit, name: unit };
						})}
						bind:value={$form.batchUnit}
					/>
				</div>
			</div>

			<!-- Gravity -->

			<div class="flex w-1/2 flex-row gap-x-4">
				<div class="flex grow flex-col gap-y-4">
					<label class="text-lg" for="originalGravity">Original Gravity</label>
					<div class="flex flex-row gap-x-4">
						<Input
							type="number"
							step="0.001"
							bind:value={$form.originalGravity}
							placeholder={$form.originalGravity.toFixed(3)}
							on:change={() => {
								const formattedNumber = Number.parseFloat($form.originalGravity.toString());
								if (formattedNumber > 0) {
									$form.originalGravity = formattedNumber;
								} else {
									$form.originalGravity = 1;
								}
							}}
						/>
					</div>
				</div>
				<div class="flex grow flex-col gap-y-4">
					<label class="text-lg" for="originalGravity">Final Gravity</label>
					<div class="flex flex-row gap-x-4">
						<Input
							type="number"
							step="0.001"
							bind:value={$form.finalGravity}
							on:change={() => {
								const formattedNumber = Number.parseFloat($form.finalGravity.toString());
								if (formattedNumber > 0) {
									$form.finalGravity = formattedNumber;
								} else {
									$form.finalGravity = 1.0;
								}
							}}
						/>
					</div>
				</div>
			</div>

			<!-- Ingredients -->

			<!-- Process Steps -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="process">Process Step</label>
				{#each $form.process as step, i}
					<div class="flex flex-row gap-x-4">
						<p class="mx-2 my-auto text-lg">{i + 1}.</p>
						<Input id="process" bind:value={$form.process[i]} />
						<Button
							variant="destructive"
							on:click={() => {
								$form.process = $form.process.filter((_, index) => index !== i);
							}}><X /></Button
						>
					</div>
				{/each}
				<Button
					on:click={() => {
						$form.process = [...$form.process, ''];
					}}
					class="mr-auto text-2xl font-bold"
				>
					+
				</Button>
			</div>

			<!-- Notes -->

			<Button type="submit">Save</Button>
		</form>
	</Card>
</div>

<DataDebug {data} />
