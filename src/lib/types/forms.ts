import { z } from 'zod';

export const emailLoginFormSchema = z.object({
	email: z.string().email(),
	password: z.string().min(8)
});
export type EmailLoginFormSchema = typeof emailLoginFormSchema;
