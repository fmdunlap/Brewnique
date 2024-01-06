import { z } from 'zod';

export const emailLoginFormSchema = z.object({
	email: z.string().email(),
	password: z.string().min(8)
});
export type EmailLoginFormSchema = typeof emailLoginFormSchema;

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

export const bioFormSchema = z.object({
	bio: z.string().max(160),
	skip: z.boolean()
});
export type BioFormSchema = typeof bioFormSchema;
