ALTER TABLE recipes ADD COLUMN author_id BIGINT NOT NULL REFERENCES users (id);

ALTER TABLE recipes ADD CONSTRAINT recipe_author_id_check CHECK (author_id IS NOT NULL AND author_id != 0);