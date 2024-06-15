CREATE TABLE IF NOT EXISTS comments (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    recipe_id BIGINT NOT NULL REFERENCES recipes (id),
    author_id BIGINT NOT NULL REFERENCES users (id),
    parent_id BIGINT REFERENCES comments (id),
    content TEXT NOT NULL
);

ALTER TABLE comments ADD CONSTRAINT comment_recipe_id_check CHECK (recipe_id IS NOT NULL AND recipe_id != 0);
ALTER TABLE comments ADD CONSTRAINT comment_author_id_check CHECK (author_id IS NOT NULL AND author_id != 0);
ALTER TABLE comments ADD CONSTRAINT comment_parent_id_check CHECK (parent_id IS NULL OR parent_id != 0);
ALTER TABLE comments ADD CONSTRAINT comment_content_check CHECK (content IS NOT NULL AND content != '');

ALTER TABLE comments ADD CONSTRAINT comment_recipe_id_unique UNIQUE (recipe_id);