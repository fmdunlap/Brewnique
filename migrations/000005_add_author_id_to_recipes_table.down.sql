ALTER TABLE recipes DROP COLUMN author_id;

ALTER TABLE recipes DROP CONSTRAINT IF EXISTS recipe_author_id_check;