import { unitOfMeasurement } from '$src/schema';
import { z } from 'zod';

export const NewRecipeFormSchema = z.object({
	id: z.string(),
	name: z
		.string()
		.min(3, { message: 'Name must be at least 3 characters.' })
		.max(100, { message: 'Name cannot be longer than 100 characters.' }),
	description: z.string().min(0).max(1000),
	images: z
		.array(z.string())
		.max(10, { message: 'Cannot have more than 10 images per post.' })
		.default([]),
	batchSize: z
		.number()
		.min(0, { message: 'Cannot have a negative batch size.' })
		.max(1000, { message: 'Batch size cannot be more than 1000. Maybe use a different unit?' })
		.default(1),
	batchUnit: z.enum(unitOfMeasurement.enumValues).default('gal'),
	originalGravity: z
		.number()
		.min(0.79, {
			message: 'Original Gravity must be more than 0.79. (Pure ethanol has a SG of 0.787)'
		})
		.max(1.75, { message: 'Original Gravity must be less than 1.75. (Theoretical max is ~1.7)' })
		.default(1.0),
	finalGravity: z
		.number()
		.min(0.79, {
			message: 'Final Gravity must be more than 0.79. (Pure ethanol has a SG of 0.787)'
		})
		.max(1.75, { message: 'Final Gravity must be less than 1.75. (Theoretical max is ~1.7)' })
		.default(1.0),
	process: z.array(z.string()).max(30, { message: 'Cannot have more than 30 steps' }).default([]),
	ingredients: z
		.array(
			z.object({
				name: z
					.string()
					.min(3, { message: 'Ingredient name must be at least 3 characters.' })
					.max(100, { message: 'Ingredient name cannot be longer than 100 characters.' }),
				quantity: z
					.number({ coerce: true })
					.min(0, { message: 'Ingredient quantity cannot be negative.' })
					.max(1000, {
						message: 'Ingredient quantity cannot be over 1000. Maybe use a different unit?'
					}),
				unit: z.enum(unitOfMeasurement.enumValues)
			})
		)
		.max(20)
		.default([]),
	notes: z
		.string()
		.min(0)
		.max(10000, { message: 'Notes cannot be longer than 10000 characters.' })
		.default('')
});
