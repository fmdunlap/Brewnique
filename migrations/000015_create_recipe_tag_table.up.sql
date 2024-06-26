CREATE TABLE IF NOT EXISTS recipe_tags (
    id BIGSERIAL PRIMARY KEY,
    recipe_id BIGINT NOT NULL REFERENCES recipes (id),
    tag_id BIGINT NOT NULL REFERENCES tags (id)
);

ALTER TABLE recipe_tags ADD CONSTRAINT recipe_tag_recipe_id_check CHECK (recipe_id != 0);
ALTER TABLE recipe_tags ADD CONSTRAINT recipe_tag_tag_id_check CHECK (tag_id != 0);