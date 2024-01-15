<script lang="ts">
	import DataDebug from '$lib/components/dev/DataDebug.svelte';
	import { Button } from '$lib/components/ui/button';
	import { Card } from '$lib/components/ui/card';
	import { Input } from '$lib/components/ui/input';
	import { superForm } from 'sveltekit-superforms/client';
	import type { PageData } from './$types';
	import Textarea from '$lib/components/ui/textarea/textarea.svelte';

	import ImageUpload from './ImageUpload.svelte';

	export let data: PageData;
	const { form, enhance } = superForm(data.form);

	let pond;
</script>

<div class="mx-auto w-3/4">
	<h1 class="py-6 text-center text-2xl">Create New Recipe</h1>
	<Card class="my-2 p-8">
		<form class="flex flex-col gap-y-12" use:enhance method="post">
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

			<ImageUpload />

			<div class="flex flex-col gap-y-4">
				<label class="text-lg" for="description">Description</label>
				<Textarea id="description" name="description" />
			</div>

			<Button type="submit">Save</Button>
		</form>
	</Card>
</div>

<DataDebug {data} />
