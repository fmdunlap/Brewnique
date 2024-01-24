import { sendEmail } from '$lib/email.js';

export const actions = {
	default: async () => {
		sendEmail('forrest@brewnique.io', 'Test', '<h1>Hello, world!</h1>', 'Hello, world!');
	}
};
