package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"time"

	"brewnique.fdunlap.com/internal/data"
	"brewnique.fdunlap.com/internal/service"
)

// TODO: A few things...
// * Function to get all comments within a comment chain
// * Pagination for list comments

type CommentAPIResponse struct {
	Id        int64                `json:"id"`
	Content   string               `json:"content"`
	Author    data.User            `json:"author"`
	RecipeId  int64                `json:"recipe_id"`
	Children  []CommentAPIResponse `json:"children"`
	CreatedAt time.Time            `json:"created_at"`
	UpdatedAt time.Time            `json:"updated_at"`
}

func (app *application) listRecipeCommentsHandler(w http.ResponseWriter, r *http.Request) {
	recipeId, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	commentData, err := app.Services.Comments.ListRecipeComments(recipeId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	comments, err := TransformComments(commentData, app.Services.Users)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	app.writeJson(w, http.StatusOK, comments, nil)
}

func TransformComments(comments []data.Comment, userService *service.UserService) ([]*CommentAPIResponse, error) {
	// Create a map to store comments by their ID for efficient lookup
	commentMap := make(map[int64]*CommentAPIResponse)

	// Create a slice to store top-level comments
	var topLevelComments []*CommentAPIResponse

	// Memoize authors during the first pass
	memoizedAuthors := make(map[int64]*data.User)

	// First pass: Create CommentAPIResponse objects and populate the commentMap
	for _, comment := range comments {
		author, ok := memoizedAuthors[comment.AuthorId]
		if !ok {
			var err error
			author, err = userService.GetUser(comment.AuthorId)
			if err != nil {
				return nil, err
			}
			memoizedAuthors[comment.AuthorId] = author
		}

		apiComment := CommentAPIResponse{
			Id:        comment.Id,
			Content:   comment.Content,
			Author:    *author,
			Children:  make([]CommentAPIResponse, 0),
			RecipeId:  comment.RecipeId,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
		commentMap[comment.Id] = &apiComment

		if comment.IsTopLevel() {
			topLevelComments = append(topLevelComments, &apiComment)
		}
	}

	// Second pass: Build the nested structure by linking children to their parent
	for _, comment := range comments {
		if comment.IsTopLevel() {
			continue
		}
		parentComment, ok := commentMap[comment.ParentId]
		if ok {
			log.Printf("Linking %v:%v to %v:%v", comment.Id, comment.Content, parentComment.Id, parentComment.Content)
			parentComment.Children = append(parentComment.Children, *commentMap[comment.Id])
		}
	}

	return topLevelComments, nil
}

func (app *application) getCommentHandler(w http.ResponseWriter, r *http.Request) {
	commentId, err := app.readIdParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
	}

	comment, err := app.Services.Comments.GetComment(commentId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, comment, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// TODO: Add route protection
func (app *application) newCommentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Content  string `json:"content"`
		AuthorId int64  `json:"author_id"`
		RecipeId int64  `json:"recipe_id"`
		ParentId int64  `json:"parent_id"`
	}

	err := app.readJson(w, r, &input)
	if err != nil {
		app.logError(r, err)
		app.badRequestResponse(w, r)
		return
	}

	comment, err := app.Services.Comments.CreateComment(input.Content, input.AuthorId, input.RecipeId, input.ParentId)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = app.writeJson(w, http.StatusOK, comment, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
