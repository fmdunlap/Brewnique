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

func (s *CommentService) CreateComment(content string, authorId, recipeId, parentId int64) (Comment, error) {
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
				return Comment{}, fmt.Errorf("parent comment %v does not exist", parentId)
			}
			return Comment{}, err
		}
		if parentComment.RecipeId != recipeId {
			return Comment{}, fmt.Errorf("parent comment is not part of recipe %v", recipeId)
		}
	}

	return s.commentProvider.PutComment(&comment)
}

func (s *CommentService) GetComment(id int64) (*Comment, error) {
	return s.commentProvider.GetComment(id)
}

func (s *CommentService) UpdateComment(comment *Comment) error {
	existingComment, err := s.commentProvider.GetComment(comment.Id)
	if err != nil {
		return err
	}
	if existingComment.RecipeId != comment.RecipeId {
		return fmt.Errorf("cannot change recipe of comment %v", comment.Id)
	}
	if existingComment.ParentId != comment.ParentId {
		return fmt.Errorf("cannot change parent of comment %v", comment.Id)
	}
	if existingComment.AuthorId != comment.AuthorId {
		return fmt.Errorf("cannot change author of comment %v", comment.Id)
	}

	return s.commentProvider.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(id int64) error {
	return s.commentProvider.DeleteComment(id)
}

func (s *CommentService) ListRecipeComments(recipeId int64) ([]Comment, error) {
	return s.commentProvider.ListRecipeComments(recipeId)
}

func (s *CommentService) ListUserComments(userId int64) ([]Comment, error) {
	return s.commentProvider.ListUserComments(userId)
}
