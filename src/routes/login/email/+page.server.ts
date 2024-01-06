import { superValidate } from 'sveltekit-superforms/server';
import { emailLoginFormSchema } from '$lib/types/forms';
import { fail } from '@sveltejs/kit';

export const load = async () => {
	return {
		form: await superValidate(emailLoginFormSchema)
	};
};

export const actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, emailLoginFormSchema);
		console.log('POST', form);

		// Convenient validation check:
		if (!form.valid) {
			// Again, return { form } and things will just work.
			return fail(400, { form });
		}

		// TODO: Do something with the validated form.data

		// Yep, return { form } here too
		return { form };
	}
};
