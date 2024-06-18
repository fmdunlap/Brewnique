CREATE TABLE comment_votes (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW(),
    user_id BIGINT NOT NULL REFERENCES users (id),
    comment_id BIGINT NOT NULL REFERENCES comments (id),
    is_upvote BOOLEAN NOT NULL DEFAULT TRUE
);

ALTER TABLE comment_votes ADD CONSTRAINT comment_vote_user_id_check CHECK (user_id != 0);
ALTER TABLE comment_votes ADD CONSTRAINT comment_vote_comment_id_check CHECK (comment_id != 0);