<script lang="ts">
	import * as Form from '$lib/components/ui/form';
	import { LoginFormSchema } from '$lib/types/forms';

	import type { SuperValidated } from 'sveltekit-superforms';

	let data: SuperValidated<typeof LoginFormSchema>;
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
			<Form.Label>Email</Form.Label>
			<Form.Input placeholder="Email" class="placeholder:text-white" />
			<Form.Validation />
		</Form.Field>
	</Form.Item>
	<Form.Item>
		<Form.Field {config} name="password">
			<Form.Label>Password</Form.Label>
			<Form.Input type="password" placeholder="Password" class="placeholder:text-white" />
			<Form.Validation />
		</Form.Field>
	</Form.Item>
	<Form.Button>Sign In</Form.Button>
</Form.Root>
