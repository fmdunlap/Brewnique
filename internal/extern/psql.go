package extern

import (
	"brewnique.fdunlap.com/internal/data"
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
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
