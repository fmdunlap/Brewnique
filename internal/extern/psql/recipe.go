package psql

import (
	"brewnique.fdunlap.com/internal/data"
	"github.com/lib/pq"
	"log"
)

func (p PostgresProvider) GetRecipe(id int64) (*data.Recipe, error) {
	res := p.db.QueryRow("SELECT id, created_at, updated_at, name, ingredients, instructions, version FROM recipes WHERE id = $1", id)

	recipe := data.Recipe{}
	err := res.Scan(&recipe)
	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (p PostgresProvider) ListRecipes() ([]*data.Recipe, error) {
	rows, err := p.db.Query("SELECT id, created_at, updated_at, name, ingredients, instructions, version FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []data.Recipe
	for rows.Next() {
		var recipe data.Recipe
		err = rows.Scan(&recipe.Id, &recipe.CreatedAt, &recipe.UpdatedAt, &recipe.Name, (*pq.StringArray)(&recipe.Ingredients), (*pq.StringArray)(&recipe.Instructions), &recipe.Version)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (p PostgresProvider) ListRecipesByAuthorId(userId int64) ([]*data.Recipe, error) {
	rows, err := p.db.Query("SELECT id, created_at, updated_at, name, ingredients, instructions, version FROM recipes WHERE author_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []*data.Recipe
	for rows.Next() {
		var recipe data.Recipe
		err = rows.Scan(
			&recipe.Id,
			&recipe.CreatedAt,
			&recipe.UpdatedAt,
			&recipe.Name,
			(*pq.StringArray)(&recipe.Ingredients),
			(*pq.StringArray)(&recipe.Instructions),
			&recipe.Version,
		)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}

func (p *PostgresProvider) PutRecipe(recipe *data.Recipe) (*data.Recipe, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	log.Printf("Creating recipe with name %s, ingredients %v, instructions %v", recipe.Name, recipe.Ingredients, recipe.Instructions)

	var insertedRecipe data.Recipe
	err = tx.QueryRow(`
        INSERT INTO recipes (name, author_id, ingredients, instructions, version)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at, author_id, name, ingredients, instructions, version
    `, recipe.Name, recipe.AuthorId, pq.Array(recipe.Ingredients), pq.Array(recipe.Instructions), recipe.Version).Scan(
		&insertedRecipe.Id,
		&insertedRecipe.CreatedAt,
		&insertedRecipe.UpdatedAt,
		&insertedRecipe.AuthorId,
		&insertedRecipe.Name,
		pq.Array(&insertedRecipe.Ingredients),
		pq.Array(&insertedRecipe.Instructions),
		&insertedRecipe.Version,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &insertedRecipe, nil
}

func (p *PostgresProvider) UpdateRecipe(recipe *data.Recipe) (*data.Recipe, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var updatedRecipe data.Recipe
	err = tx.QueryRow(`
        UPDATE recipes
        SET name = $1, ingredients = $2, instructions = $3, version = $4, updated_at = NOW()
        WHERE id = $5
        RETURNING id, created_at, updated_at, author_id, name, ingredients, instructions, version
    `, recipe.Name, pq.Array(recipe.Ingredients), pq.Array(recipe.Instructions), recipe.Version, recipe.Id).Scan(
		&updatedRecipe.Id,
		&updatedRecipe.CreatedAt,
		&updatedRecipe.UpdatedAt,
		&updatedRecipe.AuthorId,
		&updatedRecipe.Name,
		pq.Array(&updatedRecipe.Ingredients),
		pq.Array(&updatedRecipe.Instructions),
		&updatedRecipe.Version,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &updatedRecipe, nil
}

func (p PostgresProvider) DeleteRecipe(id int64) error {
	_, err := p.db.Exec("DELETE FROM recipes WHERE id = $1", id)
	return err
}
