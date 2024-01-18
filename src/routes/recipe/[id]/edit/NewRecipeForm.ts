import { unitOfMeasurement } from '$src/schema';
import { z } from 'zod';

export const NewRecipeFormSchema = z.object({
	id: z.string(),
	name: z.string().min(3).max(100),
	description: z.string().min(3).max(1000),
	images: z.array(z.string()).max(10).default([]),
	batchSize: z.number().min(1).max(1000).default(1),
	batchUnit: z.enum(unitOfMeasurement.enumValues).default('gal'),
	originalGravity: z.number().min(0.7).max(1.5).default(1.0),
	finalGravity: z.number().min(0.7).max(1.5).default(1.0),
	process: z.array(z.string()).max(20).default([]),
	ingredients: z
		.array(
			z.object({
				name: z.string().min(3).max(100),
				quantity: z.number({ coerce: true }).min(0).max(1000),
				unit: z.enum(unitOfMeasurement.enumValues)
			})
		)
		.max(20)
		.default([]),
	notes: z.string().min(3).max(1000).default('')
});
