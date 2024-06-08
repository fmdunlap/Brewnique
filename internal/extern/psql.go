package extern

import (
	"brewnique.fdunlap.com/internal/data"
	"database/sql"
)

type PsqlRecipeProvider struct {
	db *sql.DB
}

func (PsqlRecipeProvider) GetRecipe(id int64) (data.Recipe, error) {
	//TODO implement me
	panic("implement me")
}

func (PsqlRecipeProvider) ListRecipes() ([]data.Recipe, error) {
	//TODO implement me
	panic("implement me")
}

func (PsqlRecipeProvider) CreateRecipe(recipe data.Recipe) (data.Recipe, error) {
	//TODO implement me
	panic("implement me")
}

func (PsqlRecipeProvider) UpdateRecipe(recipe data.Recipe) (data.Recipe, error) {
	//TODO implement me
	panic("implement me")
}

func (PsqlRecipeProvider) DeleteRecipe(id int64) error {
	//TODO implement me
	panic("implement me")
}
