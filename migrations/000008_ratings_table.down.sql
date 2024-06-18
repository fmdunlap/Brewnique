DROP TABLE IF EXISTS ratings;

ALTER TABLE comments DROP CONSTRAINT IF EXISTS rating_user_id_check;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS rating_recipe_id_check;