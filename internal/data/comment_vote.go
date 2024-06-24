package data

import "time"

type CommentVote struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CommentId int64     `json:"comment_id"`
	UserId    int64     `json:"user_id"`
	IsUpVote  bool      `json:"is_upvote"`
}
