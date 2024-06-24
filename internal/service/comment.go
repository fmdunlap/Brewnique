package service

import (
	"brewnique.fdunlap.com/internal/data"
	"database/sql"
	"errors"
	"fmt"
)

type CommentProvider interface {
	AddCommentVote(commentId int64, userId int64, isUpVote bool) error
	DeleteComment(id int64) error
	DeleteCommentVote(commentId int64, userId int64) error
	GetCommentVotes(commentId int64) ([]*data.CommentVote, error)
	GetComment(id int64) (*data.Comment, error)
	ListRecipeComments(recipeId int64) ([]data.Comment, error)
	ListUserComments(userId int64) ([]data.Comment, error)
	PutComment(comment *data.Comment) (data.Comment, error)
	UpdateComment(comment *data.Comment) error
}

type CommentService struct {
	commentProvider CommentProvider
}

func NewCommentService(commentProvider CommentProvider) *CommentService {
	return &CommentService{
		commentProvider: commentProvider,
	}
}

func (s *CommentService) CreateComment(content string, authorId, recipeId, parentId int64) (*data.Comment, error) {
	comment := data.Comment{
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

func (s *CommentService) GetComment(id int64) (*data.Comment, error) {
	return s.commentProvider.GetComment(id)
}

func (s *CommentService) UpdateComment(id int64, content string) (*data.Comment, error) {
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

func (s *CommentService) ListRecipeComments(recipeId int64) ([]data.Comment, error) {
	if recipeId == 0 {
		return nil, fmt.Errorf("recipeId is not set")
	}
	comments, err := s.commentProvider.ListRecipeComments(recipeId)
	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return make([]data.Comment, 0), nil
	}
	return comments, nil
}

func (s *CommentService) ListUserComments(userId int64) ([]data.Comment, error) {
	return s.commentProvider.ListUserComments(userId)
}

func (s *CommentService) GetCommentScore(commentId int64) (int64, error) {
	votes, err := s.commentProvider.GetCommentVotes(commentId)
	if err != nil {
		return 0, err
	}

	score := int64(0)
	for _, vote := range votes {
		if vote.IsUpVote {
			score++
		} else {
			score--
		}
	}

	return score, nil
}

func (s *CommentService) UpvoteComment(commentId int64, userId int64) error {
	return s.commentProvider.AddCommentVote(commentId, userId, true)
}

func (s *CommentService) DownvoteComment(commentId int64, userId int64) error {
	return s.commentProvider.AddCommentVote(commentId, userId, false)
}
