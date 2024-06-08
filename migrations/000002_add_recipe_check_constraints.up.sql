ALTER TABLE recipes ADD CONSTRAINT recipe_name_check CHECK (name IS NOT NULL AND name != '');
ALTER TABLE recipes ADD CONSTRAINT recipe_ingredients_check_constraint CHECK (array_length(ingredients, 1) > 0 AND ingredients IS NOT NULL);
ALTER TABLE recipes ADD CONSTRAINT recipe_instructions_check_constraint CHECK (array_length(instructions, 1) > 0 AND instructions IS NOT NULL);