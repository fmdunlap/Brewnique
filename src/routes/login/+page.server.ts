import { fail, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import { superValidate } from 'sveltekit-superforms/server';
import { LoginFormSchema } from '$lib/types/forms';

export const load: PageServerLoad = async () => {
	return {
		form: await superValidate(LoginFormSchema)
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
		return {
			form
		};
	}
};
