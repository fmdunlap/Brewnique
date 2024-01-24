import { z } from 'zod';

const DisplayNameRegexp = new RegExp('^[a-z0-9_-]{5,20}$');

export const usernameFormSchema = z.object({
	username: z
		.string()
		.regex(
			DisplayNameRegexp,
			"Display name can only have lowercase alphabetical characters, numbers, and '_' underscores."
		)
		.min(5)
		.max(20)
});
export type UsernameFormSchema = typeof usernameFormSchema;
