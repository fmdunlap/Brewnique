import { createInsertSchema } from 'drizzle-zod';
import { recipe } from '$src/schema';

export const NewRecipeFormSchema = createInsertSchema(recipe);
