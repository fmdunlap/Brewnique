package psql

import (
	"brewnique.fdunlap.com/internal/data"
	"database/sql"
	"errors"
)

func (p *PostgresProvider) PutComment(comment *data.Comment) (data.Comment, error) {
	var existingID int64
	existingComment, err := p.GetComment(comment.Id)
	if err == nil {
		existingID = existingComment.Id
	} else if !errors.Is(err, sql.ErrNoRows) {
		return data.Comment{}, err
	}

	if existingID > 0 {
		return data.Comment{}, errors.New("comment already exists")
	}

	tx, err := p.db.Begin()
	if err != nil {
		return data.Comment{}, err
	}
	defer tx.Rollback()

	err = tx.QueryRow("INSERT INTO comments (content, author_id, recipe_id, parent_id) VALUES ($1, $2, $3, $4) RETURNING id", comment.Content, comment.AuthorId, comment.RecipeId, comment.ParentId).Scan(&comment.Id)
	if err != nil {
		return data.Comment{}, err
	}

	err = tx.Commit()
	if err != nil {
		return data.Comment{}, err
	}

	return *comment, nil
}

func (p *PostgresProvider) GetComment(id int64) (*data.Comment, error) {
	comment := data.Comment{}
	err := p.db.QueryRow("SELECT id, created_at, updated_at, recipe_id, author_id, parent_id, content FROM comments WHERE id = $1", id).Scan(&comment.Id, &comment.CreatedAt, &comment.UpdatedAt, &comment.RecipeId, &comment.AuthorId, &comment.ParentId, &comment.Content)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (p *PostgresProvider) UpdateComment(comment *data.Comment) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if comment.Content == "" {
		return errors.New("content cannot be empty")
	}

	tx.QueryRow("UPDATE comments SET content = $1, updated_at = NOW() WHERE id = $2", comment.Content, comment.Id)
	return tx.Commit()
}

func (p *PostgresProvider) ListRecipeComments(recipeId int64) ([]data.Comment, error) {
	comments := []data.Comment{}
	rows, err := p.db.Query("SELECT id, created_at, updated_at, recipe_id, author_id, parent_id, content FROM comments WHERE recipe_id = $1", recipeId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment data.Comment
		err = rows.Scan(&comment.Id, &comment.CreatedAt, &comment.UpdatedAt, &comment.RecipeId, &comment.AuthorId, &comment.ParentId, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (p *PostgresProvider) ListUserComments(userId int64) ([]data.Comment, error) {
	comments := []data.Comment{}
	rows, err := p.db.Query("SELECT id, created_at, updated_at, recipe_id, author_id, parent_id, content FROM comments WHERE author_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment data.Comment
		err = rows.Scan(&comment.Id, &comment.CreatedAt, &comment.UpdatedAt, &comment.RecipeId, &comment.AuthorId, &comment.ParentId, &comment.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (p *PostgresProvider) DeleteComment(id int64) error {
	_, err := p.db.Exec("DELETE FROM comments WHERE id = $1", id)
	return err
}
