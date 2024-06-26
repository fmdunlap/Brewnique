CREATE TABLE IF NOT EXISTS recipe_attributes (
    id BIGSERIAL PRIMARY KEY,
    recipe_id BIGINT NOT NULL REFERENCES recipes (id),
    attribute_id BIGINT NOT NULL REFERENCES attributes (id)
);

ALTER TABLE recipe_attributes ADD CONSTRAINT recipe_attribute_recipe_id_check CHECK (recipe_id != 0);
ALTER TABLE recipe_attributes ADD CONSTRAINT recipe_attribute_attribute_id_check CHECK (attribute_id != 0);