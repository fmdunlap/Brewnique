import { fail, type Actions } from '@sveltejs/kit';
import { superValidate } from 'sveltekit-superforms/server';
import { LoginFormSchema } from '$lib/types/forms';

export const load = async (event) => {
	return {
		form: await superValidate(event, LoginFormSchema)
	};
};

export const actions: Actions = {
	default: async (event) => {
		const form = await superValidate(event, LoginFormSchema);
		if (!form.valid) {
			return fail(400, {
				form
			});
		}

		// if user has not created an account...

		// if user has created an account but is not verified...

		// if user has created an account and is verified...

		return {
			form
		};
	}
};
