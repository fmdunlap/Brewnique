ALTER TABLE recipes DROP CONSTRAINT IF EXISTS recipe_name_check;
ALTER TABLE recipes DROP CONSTRAINT IF EXISTS recipe_ingredients_check_constraint;
ALTER TABLE recipes DROP CONSTRAINT IF EXISTS recipe_instructions_check_constraint;