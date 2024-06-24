package psql

import (
	"brewnique.fdunlap.com/internal/data"
	"database/sql"
	"errors"
	"log"
	"time"
)

type CommentDbRow struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	RecipeId  int64
	AuthorId  int64
	ParentId  sql.NullInt64
	Content   string
}

func (c *CommentDbRow) ToComment() data.Comment {
	return data.Comment{
		Id:        c.Id,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		RecipeId:  c.RecipeId,
		AuthorId:  c.AuthorId,
		ParentId:  c.ParentId.Int64,
		Content:   c.Content,
	}
}

func CommentRowFromComment(comment data.Comment) CommentDbRow {
	parentId := sql.NullInt64{Int64: comment.ParentId, Valid: comment.ParentId != 0}

	return CommentDbRow{
		Id:        comment.Id,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		RecipeId:  comment.RecipeId,
		AuthorId:  comment.AuthorId,
		ParentId:  parentId,
		Content:   comment.Content,
	}
}

type CommentVoteDbRow struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	CommentId int64
	UserId    int64
	IsUpVote  bool
}

func (c *CommentVoteDbRow) ToCommentVote() data.CommentVote {
	return data.CommentVote{
		Id:        c.Id,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
		CommentId: c.CommentId,
		UserId:    c.UserId,
		IsUpVote:  c.IsUpVote,
	}
}

func (p PostgresProvider) PutComment(comment *data.Comment) (data.Comment, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return data.Comment{}, err
	}
	defer tx.Rollback()

	var insertedCommentRow CommentDbRow
	insertableComment := CommentRowFromComment(*comment)

	log.Printf("inserting comment %v", insertableComment)

	err = tx.QueryRow("INSERT INTO comments (content, author_id, recipe_id, parent_id) VALUES ($1, $2, $3, $4) RETURNING id, content, author_id, recipe_id, parent_id, created_at, updated_at",
		insertableComment.Content,
		insertableComment.AuthorId,
		insertableComment.RecipeId,
		insertableComment.ParentId,
	).Scan(
		&insertedCommentRow.Id,
		&insertedCommentRow.Content,
		&insertedCommentRow.AuthorId,
		&insertedCommentRow.RecipeId,
		&insertedCommentRow.ParentId,
		&insertedCommentRow.CreatedAt,
		&insertedCommentRow.UpdatedAt,
	)
	if err != nil {
		return data.Comment{}, err
	}

	err = tx.Commit()
	if err != nil {
		return data.Comment{}, err
	}

	insertedComment := insertedCommentRow.ToComment()
	return insertedComment, nil
}

func (p PostgresProvider) GetComment(id int64) (*data.Comment, error) {
	commentRow := CommentDbRow{}
	err := p.db.QueryRow("SELECT id, created_at, updated_at, recipe_id, author_id, parent_id, content FROM comments WHERE id = $1", id).Scan(&commentRow.Id, &commentRow.CreatedAt, &commentRow.UpdatedAt, &commentRow.RecipeId, &commentRow.AuthorId, &commentRow.ParentId, &commentRow.Content)
	if err != nil {
		return nil, err
	}

	comment := commentRow.ToComment()
	return &comment, nil
}

func (p PostgresProvider) UpdateComment(comment *data.Comment) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	tx.QueryRow("UPDATE comments SET content = $1, updated_at = NOW() WHERE id = $2", comment.Content, comment.Id)
	return tx.Commit()
}

func (p PostgresProvider) ListRecipeComments(recipeId int64) ([]data.Comment, error) {
	commentRows := []CommentDbRow{}
	rows, err := p.db.Query("SELECT id, created_at, updated_at, recipe_id, author_id, parent_id, content FROM comments WHERE recipe_id = $1", recipeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commentRow CommentDbRow
		err = rows.Scan(
			&commentRow.Id,
			&commentRow.CreatedAt,
			&commentRow.UpdatedAt,
			&commentRow.RecipeId,
			&commentRow.AuthorId,
			&commentRow.ParentId,
			&commentRow.Content,
		)
		if err != nil {
			return nil, err
		}
		commentRows = append(commentRows, commentRow)
	}

	comments := make([]data.Comment, len(commentRows))
	for i, commentRow := range commentRows {
		comments[i] = commentRow.ToComment()
	}

	return comments, nil
}

func (p PostgresProvider) ListUserComments(userId int64) ([]data.Comment, error) {
	commentDbRows := []CommentDbRow{}
	rows, err := p.db.Query("SELECT id, created_at, updated_at, recipe_id, author_id, parent_id, content FROM comments WHERE author_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var commentRow CommentDbRow
		err = rows.Scan(
			&commentRow.Id,
			&commentRow.CreatedAt,
			&commentRow.UpdatedAt,
			&commentRow.RecipeId,
			&commentRow.AuthorId,
			&commentRow.ParentId,
			&commentRow.Content,
		)
		if err != nil {
			return nil, err
		}
		commentDbRows = append(commentDbRows, commentRow)
	}

	comments := make([]data.Comment, len(commentDbRows))
	for i, commentRow := range commentDbRows {
		comments[i] = commentRow.ToComment()
	}

	return comments, nil
}

func (p PostgresProvider) DeleteComment(id int64) error {
	_, err := p.db.Exec("DELETE FROM comments WHERE id = $1", id)
	return err
}

func (p PostgresProvider) GetCommentVote(commentId int64, userId int64) (*data.CommentVote, error) {
	voteRow := CommentVoteDbRow{}
	err := p.db.QueryRow("SELECT id, created_at, updated_at, comment_id, user_id, is_upvote FROM comment_votes WHERE comment_id = $1 AND user_id = $2", commentId, userId).Scan(
		&voteRow.Id,
		&voteRow.CreatedAt,
		&voteRow.UpdatedAt,
		&voteRow.CommentId,
		&voteRow.UserId,
		&voteRow.IsUpVote,
	)
	if err != nil {
		return &data.CommentVote{}, err
	}

	vote := voteRow.ToCommentVote()

	return &vote, nil
}

func (p PostgresProvider) GetCommentVotes(commentId int64) ([]*data.CommentVote, error) {
	votes := []*data.CommentVote{}
	rows, err := p.db.Query("SELECT id, created_at, updated_at, comment_id, user_id, is_upvote FROM comment_votes WHERE comment_id = $1", commentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var voteRow CommentVoteDbRow
		err = rows.Scan(
			&voteRow.Id,
			&voteRow.CreatedAt,
			&voteRow.UpdatedAt,
			&voteRow.CommentId,
			&voteRow.UserId,
			&voteRow.IsUpVote,
		)
		if err != nil {
			return nil, err
		}

		vote := voteRow.ToCommentVote()
		votes = append(votes, &vote)
	}

	return votes, nil
}

func (p PostgresProvider) AddCommentVote(commentId int64, userId int64, isUpVote bool) error {
	// Check if the user already voted on this comment
	vote, err := p.GetCommentVote(commentId, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if vote != nil {
		if vote.IsUpVote == isUpVote {
			return nil
		}
		err = p.DeleteCommentVote(commentId, userId)
		if err != nil {
			return err
		}
	}

	_, err = p.db.Exec("INSERT INTO comment_votes (comment_id, user_id, is_upvote) VALUES ($1, $2, $3)", commentId, userId, isUpVote)
	return err
}

func (p PostgresProvider) DeleteCommentVote(commentId int64, userId int64) error {
	_, err := p.db.Exec("DELETE FROM comment_votes WHERE comment_id = $1 AND user_id = $2", commentId, userId)
	return err
}
