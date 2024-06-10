package extern

import (
	"brewnique.fdunlap.com/internal/data"
	"context"
	"database/sql"
	"github.com/lib/pq"
	"log"
	"time"
)

type PsqlConfig struct {
	Dsn             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type PsqlRecipeProvider struct {
	db *sql.DB
}

func NewPsqlProvider(config PsqlConfig) *PsqlRecipeProvider {
	db, err := sql.Open("postgres", config.Dsn)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	log.Printf("Connected to database")

	return &PsqlRecipeProvider{
		db: db,
	}
}

func (p *PsqlRecipeProvider) Close() {
	p.db.Close()
}

func (p PsqlRecipeProvider) GetRecipe(id int64) (data.Recipe, error) {
	rows, err := p.db.Query("SELECT id, created_at, updated_at, name, ingredients, instructions, version FROM recipes WHERE id = $1", id)
	if err != nil {
		return data.Recipe{}, err
	}
	defer rows.Close()

	var recipe data.Recipe
	for rows.Next() {
		err = rows.Scan(&recipe.ID, &recipe.CreatedAt, &recipe.UpdatedAt, &recipe.Name, (*pq.StringArray)(&recipe.Ingredients), (*pq.StringArray)(&recipe.Instructions), &recipe.Version)
		if err != nil {
			return data.Recipe{}, err
		}
	}

	return recipe, nil
}

func (p PsqlRecipeProvider) ListRecipes() ([]data.Recipe, error) {
	rows, err := p.db.Query("SELECT id, created_at, updated_at, name, ingredients, instructions, version FROM recipes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var recipes []data.Recipe
	for rows.Next() {
		var recipe data.Recipe
		err = rows.Scan(&recipe.ID, &recipe.CreatedAt, &recipe.UpdatedAt, &recipe.Name, (*pq.StringArray)(&recipe.Ingredients), (*pq.StringArray)(&recipe.Instructions), &recipe.Version)
		if err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

func (p PsqlRecipeProvider) CreateRecipe(recipe data.Recipe) (data.Recipe, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return data.Recipe{}, err
	}

	log.Printf("Creating recipe with name %s, ingredients %v, instructions %v", recipe.Name, recipe.Ingredients, recipe.Instructions)

	var id int64
	err = tx.QueryRow("INSERT INTO recipes (name, ingredients, instructions, version) VALUES ($1, $2, $3, $4) RETURNING id", recipe.Name, pq.Array(recipe.Ingredients), pq.Array(recipe.Instructions), recipe.Version).Scan(&id)
	if err != nil {
		tx.Rollback()
		return data.Recipe{}, err
	}

	err = tx.Commit()
	if err != nil {
		return data.Recipe{}, err
	}

	recipe.ID = id

	return recipe, nil
}

func (p PsqlRecipeProvider) UpdateRecipe(recipe data.Recipe) (data.Recipe, error) {
	tx, err := p.db.Begin()

	if err != nil {
		return data.Recipe{}, err
	}

	tx.QueryRow("UPDATE recipes SET name = $1, ingredients = $2, instructions = $3, version = $4 WHERE id = $5", recipe.Name, pq.Array(recipe.Ingredients), pq.Array(recipe.Instructions), recipe.Version, recipe.ID)
	err = tx.Commit()
	if err != nil {
		return data.Recipe{}, err
	}

	return recipe, nil
}

func (p PsqlRecipeProvider) DeleteRecipe(id int64) error {
	_, err := p.db.Exec("DELETE FROM recipes WHERE id = $1", id)
	return err
}
