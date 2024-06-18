CREATE TABLE ratings (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    user_id BIGINT NOT NULL REFERENCES users (id),
    recipe_id BIGINT NOT NULL REFERENCES recipes (id),
    rating INTEGER NOT NULL
);

ALTER TABLE ratings ADD CONSTRAINT rating_user_id_check CHECK (user_id != 0);
ALTER TABLE ratings ADD CONSTRAINT rating_recipe_id_check CHECK (recipe_id != 0);