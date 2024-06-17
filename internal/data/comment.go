package data

import (
	"brewnique.fdunlap.com/internal/validator"
	"database/sql"
	"errors"
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

type CommentProvider interface {
	PutComment(comment *Comment) (Comment, error)
	GetComment(id int64) (*Comment, error)
	UpdateComment(comment *Comment) error
	ListRecipeComments(recipeId int64) ([]Comment, error)
	ListUserComments(userId int64) ([]Comment, error)
	DeleteComment(id int64) error
}

type CommentService struct {
	commentProvider CommentProvider
}

func NewCommentService(commentProvider CommentProvider) *CommentService {
	return &CommentService{
		commentProvider: commentProvider,
	}
}

func (s *CommentService) CreateComment(content string, authorId, recipeId, parentId int64) (*Comment, error) {
	comment := Comment{
		Content:  content,
		AuthorId: authorId,
		RecipeId: recipeId,
		ParentId: parentId,
	}

	if parentId > 0 {
		parentComment, err := s.commentProvider.GetComment(parentId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, fmt.Errorf("parent comment %v does not exist", parentId)
			}
			return nil, err
		}
		if parentComment.RecipeId != recipeId {
			return nil, fmt.Errorf("parent comment is not part of recipe %v", recipeId)
		}
	}

	if authorId == 0 {
		return nil, fmt.Errorf("authorId is not set")
	}
	if recipeId == 0 {
		return nil, fmt.Errorf("recipeId is not set")
	}

	comment, err := s.commentProvider.PutComment(&comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (s *CommentService) GetComment(id int64) (*Comment, error) {
	return s.commentProvider.GetComment(id)
}

func (s *CommentService) UpdateComment(id int64, content string) (*Comment, error) {
	comment, err := s.commentProvider.GetComment(id)
	if err != nil {
		return nil, err
	}

	comment.Content = content

	err = s.commentProvider.UpdateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *CommentService) DeleteComment(id int64) error {
	return s.commentProvider.DeleteComment(id)
}

func (s *CommentService) ListRecipeComments(recipeId int64) ([]Comment, error) {
	if recipeId == 0 {
		return nil, fmt.Errorf("recipeId is not set")
	}
	comments, err := s.commentProvider.ListRecipeComments(recipeId)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return make([]Comment, 0), nil
	}
	return comments, nil
}

func (s *CommentService) ListUserComments(userId int64) ([]Comment, error) {
	return s.commentProvider.ListUserComments(userId)
}
