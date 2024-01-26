import { z } from 'zod';

// const UsernameRegexp = new RegExp('^[a-z0-9_-]{5,20}$');

export const usernameFormSchema = z.object({
	username: z.string().min(5).max(20)
});
export type UsernameFormSchema = typeof usernameFormSchema;
