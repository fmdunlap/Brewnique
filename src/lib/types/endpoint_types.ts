import type { Enums } from './supabaseDB';

export interface RecipeIngredient {
	name: string;
	type: Enums<'ingredient_type'>;
	quantity: number;
	unit: Enums<'unit_of_measurement'>;
}

export interface Recipe {
	id: number;
	user_id: number;
	name: string;
	description: string | null;
	published: boolean;
	difficulty: Enums<'difficulty'>;
	type: Enums<'brew_type'>;
	original_gravity: number | null;
	final_gravity: number | null;
	sweetened_gravity: number | null;
	process_steps: string[] | null;
	notes: string | null;
	ingredients: RecipeIngredient[];
}
