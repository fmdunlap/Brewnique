CREATE TABLE IF NOT EXISTS users {
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    username TEXT NOT NULL,
    email TEXT NOT NULL,
};

ALTER TABLE users ADD CONSTRAINT user_username_check CHECK (username IS NOT NULL AND username != '');
ALTER TABLE users ADD CONSTRAINT user_email_check CHECK (email IS NOT NULL AND email != '');

ALTER TABLE users ADD CONSTRAINT user_email_unique UNIQUE (email);
ALTER TABLE users ADD CONSTRAINT user_username_unique UNIQUE (username);