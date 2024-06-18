DROP TABLE IF EXISTS comment_votes;

ALTER TABLE comments DROP CONSTRAINT IF EXISTS comment_vote_user_id_check;
ALTER TABLE comments DROP CONSTRAINT IF EXISTS comment_vote_comment_id_check;