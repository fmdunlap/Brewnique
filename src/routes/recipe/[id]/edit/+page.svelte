<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { Card } from '$lib/components/ui/card';
	import { Button } from '$lib/components/ui/button';
	import { Input, Textarea, Select } from 'flowbite-svelte';
	import { superForm } from 'sveltekit-superforms/client';

	import ImageUpload from './ImageUpload.svelte';
	import type { PageData } from './$types';
	import { Plus, Trash2 } from 'lucide-svelte';
	import { recipeIngredient } from '$src/schema';
	import FormError from '$lib/components/form/FormError.svelte';

	export let data: PageData;
	const { form, enhance, errors } = superForm(data.form, {
		dataType: 'json'
	});

	const batchSizeUnits = ['gal', 'liter', 'barrel', 'cup', 'oz'];
</script>

<div class="mx-auto w-3/4">
	<h1 class="py-6 text-center text-2xl">Create New Recipe</h1>
	<Card class="my-2 p-8">
		<form class="flex flex-col gap-y-12" action="?/save" use:enhance method="post">
			<!-- Hidden ID -->

			<input type="hidden" name="id" bind:value={$form.id} />

			<!-- Name -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="name">Name</label>
				<Input
					type="text"
					name="name"
					placeholder="The marvelous melon mead"
					bind:value={$form.name}
				/>
				<FormError errorMessages={$errors.name} />
			</div>

			<!-- Images -->

			<ImageUpload bind:b64imgs={$form.images} />

			{#each $form.images as _, i}
				<input type="hidden" name="images" bind:value={$form.images[i]} />
			{/each}

			<!-- Description -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="description">Description</label>
				<Textarea id="description" name="description" bind:value={$form.description} />
			</div>

			<!-- Batch Size -->

			<div class="flex w-1/2 flex-col gap-y-4">
				<label class="text-lg" for="batchSize">Batch Size</label>
				<div class="flex flex-row gap-x-4">
					<Input
						type="number"
						name="batchSize"
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
						name="batchUnit"
						items={batchSizeUnits.map((unit) => {
							return { value: unit, name: unit };
						})}
						bind:value={$form.batchUnit}
					/>
				</div>
				<FormError errorMessages={$errors.batchSize} />
				<FormError errorMessages={$errors.batchUnit} />
			</div>

			<!-- Gravity -->

			<div class="flex w-1/2 flex-row gap-x-4">
				<div class="flex grow flex-col gap-y-4">
					<label class="text-lg" for="originalGravity">Original Gravity</label>
					<div class="flex flex-row gap-x-4">
						<Input
							type="number"
							name="originalGravity"
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
					<FormError errorMessages={$errors.originalGravity} />
				</div>
				<div class="flex grow flex-col gap-y-4">
					<label class="text-lg" for="originalGravity">Final Gravity</label>
					<div class="flex flex-row gap-x-4">
						<Input
							type="number"
							name="finalGravity"
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
					<FormError errorMessages={$errors.finalGravity} />
				</div>
			</div>

			<!-- Ingredients -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="ingredients">Ingredients</label>
				{#each $form.ingredients as _, i}
					<div class="flex flex-row gap-x-4">
						<p class="mx-2 my-auto text-lg">{i + 1}.</p>
						<Input
							class="w-1/3"
							type="number"
							name="quantity"
							step="0.1"
							bind:value={$form.ingredients[i].quantity}
						/>
						<Select
							class="w-1/3"
							name="unit"
							items={recipeIngredient.unit.enumValues.map((unit) => {
								return { value: unit, name: unit };
							})}
							bind:value={$form.ingredients[i].unit}
						/>
						<Input
							name="name"
							placeholder="Ingredient Name"
							bind:value={$form.ingredients[i].name}
						/>
						<Button
							variant="destructive"
							class="p-2"
							on:click={() => {
								$form.ingredients = $form.ingredients.filter((_, index) => index !== i);
							}}
						>
							<Trash2 />
						</Button>
					</div>
					{#if $errors.ingredients && $errors.ingredients[i]}
						<FormError errorMessages={$errors.ingredients[i].name} />
						<FormError errorMessages={$errors.ingredients[i].quantity} />
						<FormError errorMessages={$errors.ingredients[i].unit} />
					{/if}
				{/each}
				<Button
					on:click={() => {
						$form.ingredients = [...$form.ingredients, { name: '', quantity: 0, unit: 'lb' }];
					}}
					class="mr-auto p-2 text-2xl font-bold"
				>
					<Plus />
				</Button>
			</div>

			<!-- Process Steps -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="process">Process Step</label>
				{#each $form.process as _, i}
					<div class="flex flex-row gap-x-4">
						<p class="mx-2 my-auto text-lg">{i + 1}.</p>
						<Input id="process" name="process" bind:value={$form.process[i]} />
						<Button
							variant="destructive"
							class="p-2"
							on:click={() => {
								$form.process = $form.process.filter((_, index) => index !== i);
							}}><Trash2 /></Button
						>
					</div>
				{/each}
				<Button
					on:click={() => {
						$form.process = [...$form.process, ''];
					}}
					class="mr-auto p-2 text-2xl font-bold"
				>
					<Plus />
				</Button>
			</div>

			<!-- Notes -->

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="notes">Notes</label>
				<Textarea name="notes" bind:value={$form.notes} />
			</div>

			<div class="flex w-full flex-col gap-y-2">
				<Button type="submit" variant="secondary">Save</Button>
				<Button type="submit" formaction="?/publish">Publish</Button>
			</div>
		</form>
	</Card>
</div>

<DataDebug {data} />
