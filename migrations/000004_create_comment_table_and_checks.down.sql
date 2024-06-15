ALTER TABLE comments DROP CONSTRAINT IF EXISTS comment_recipe_id_check;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comment_author_id_check;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comment_parent_id_check;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comment_content_check;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comment_recipe_id_unique;
DROP TABLE IF EXISTS comments;
