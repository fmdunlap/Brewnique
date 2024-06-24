package data

import (
	"brewnique.fdunlap.com/internal/validator"
	"fmt"
	"time"
)

const (
	MaxCommentContentLength = 2048
)

type Comment struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	AuthorId  int64     `json:"author_id"`
	RecipeId  int64     `json:"recipe_id"`
	ParentId  int64     `json:"parent_id"`
	Content   string    `json:"content"`
}

type CommentVote struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CommentId int64     `json:"comment_id"`
	UserId    int64     `json:"user_id"`
	IsUpVote  bool      `json:"is_upvote"`
}

func (c *Comment) String() string {
	return fmt.Sprintf("Comment{Id: %d, Content: %s, AuthorId: %d, RecipeId: %d, ParentId: %d}", c.Id, c.Content, c.AuthorId, c.RecipeId, c.ParentId)
}

func (c *Comment) IsTopLevel() bool {
	return c.ParentId == 0
}

func (c *Comment) Equal(other *Comment) bool {
	if c == nil && other == nil {
		return true
	}
	if c == nil || other == nil {
		return false
	}
	return c.Id == other.Id &&
		c.Content == other.Content &&
		c.AuthorId == other.AuthorId &&
		c.RecipeId == other.RecipeId &&
		c.ParentId == other.ParentId

}

func ValidateComment(v *validator.Validator, comment Comment) {
	v.Check(len(comment.Content) > 0, "content", "content is required")
	v.Check(len(comment.Content) < MaxCommentContentLength, "content", "content is too long")
}
