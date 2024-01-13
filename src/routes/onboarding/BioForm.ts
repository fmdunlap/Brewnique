import { z } from 'zod';

export const bioFormSchema = z.object({
	bio: z.string().max(160),
	skip: z.boolean()
});
export type BioFormSchema = typeof bioFormSchema;
