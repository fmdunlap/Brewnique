import { unitOfMeasurement } from '$src/schema';
import { z } from 'zod';

export const NewRecipeFormSchema = z.object({
	name: z.string().min(3).max(100),
	description: z.string().min(3).max(1000),
	images: z.array(z.string()).max(10),
	batchQuantity: z.number().min(1).max(1000).default(1),
	batchUnit: z.enum(unitOfMeasurement.enumValues).default('gal')
});
