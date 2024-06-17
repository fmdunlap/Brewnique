package main

import (
	"brewnique.fdunlap.com/internal/data"
	"brewnique.fdunlap.com/internal/extern/psql"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.1"

type Services struct {
	Recipes  *data.RecipeService
	Users    *data.UserService
	Comments *data.CommentService
}

type application struct {
	config   applicationConfig
	logger   *log.Logger
	Services Services
}

func main() {
	cfg := loadConfig()

	logger := log.New(os.Stdout, "", log.LstdFlags)
	dbProvider := psql.NewPsqlProvider(psql.PsqlConfig{
		Dsn:             cfg.database.dsn,
		MaxOpenConns:    cfg.database.maxOpenConns,
		MaxIdleConns:    cfg.database.maxIdleConns,
		ConnMaxLifetime: cfg.database.connMaxLifetime,
	})
	defer dbProvider.Close()

	app := &application{
		config: cfg,
		logger: logger,
		Services: Services{
			Recipes:  data.NewRecipeService(dbProvider),
			Users:    data.NewUserService(dbProvider),
			Comments: data.NewCommentService(dbProvider),
		},
	}

	log.Printf("Starting %s API server on port %d", app.config.env, app.config.port)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  cfg.http.idleTimeout,
		ReadTimeout:  cfg.http.readTimeout,
		WriteTimeout: cfg.http.writeTimeout,
	}

	err := server.ListenAndServe()
	logger.Fatal(err)
}
