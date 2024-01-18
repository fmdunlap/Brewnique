import { z } from 'zod';

const DisplayNameRegexp = new RegExp('^[a-z0-9_-]{5,20}$');

export const displayNameFormSchema = z.object({
	display_name: z
		.string()
		.regex(
			DisplayNameRegexp,
			"Display name can only have lowercase alphabetical characters, numbers, and '_' underscores."
		)
		.min(5)
		.max(20)
});
export type DisplayNameFormSchema = typeof displayNameFormSchema;
