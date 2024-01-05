import { z } from 'zod';

export const LoginFormSchema = z.object({
	email: z.string().email()
});
