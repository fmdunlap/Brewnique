import { superValidate } from 'sveltekit-superforms/server';
import { emailSignupFormSchema } from './EmailSignupFormSchema';
import { fail } from '@sveltejs/kit';

export const load = async () => {
	return {
		form: await superValidate(emailSignupFormSchema)
	};
};

export const actions = {
	default: async ({ request }) => {
		const form = await superValidate(request, emailSignupFormSchema);

		// Convenient validation check:
		if (!form.valid) {
			if (form.errors._errors) {
				form.errors.password = form.errors._errors;
			}
			// Clear the password fields
			form.data.password = '';
			form.data.confirmPassword = '';
			return fail(400, { form });
		}

		// Custom validation check for password confirmation:

		if (form.data.password !== form.data.confirmPassword) {
			if (!form.errors) form.errors = {};
			if (!form.errors.confirmPassword) form.errors.confirmPassword = [];
			form.errors.confirmPassword.push('Passwords do not match');
			// Clear the password fields
			form.data.password = '';
			form.data.confirmPassword = '';
			return fail(400, { form });
		}

		// Create the new user

		// Yep, return { form } here too
		return { form };
	}
};
