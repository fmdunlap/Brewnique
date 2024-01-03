<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { LoginFormSchema } from '$lib/types/forms';

	import type { SuperValidated } from 'sveltekit-superforms';

	export let data: SuperValidated<typeof LoginFormSchema>;
</script>

<Form.Root
	form={data}
	schema={LoginFormSchema}
	let:config
	method="POST"
	action="/auth/sign-in"
	class="mx-auto flex h-fit w-full flex-col"
>
	<Form.Item>
		<Form.Field {config} name="email">
			<Form.Input placeholder="Email" class="bg-white placeholder:text-orange-700" />
			<Form.Validation />
		</Form.Field>
	</Form.Item>
	<Form.Button>Continue with Email</Form.Button>
</Form.Root>
