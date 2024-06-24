package main

import (
	"brewnique.fdunlap.com/internal/data/service"
	"brewnique.fdunlap.com/internal/extern/psql"
	"brewnique.fdunlap.com/internal/jsonlog"
	service2 "brewnique.fdunlap.com/internal/service"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.1"

type Services struct {
	Recipes  *service.RecipeService
	Users    *service2.UserService
	Comments *service2.CommentService
}

type application struct {
	config   applicationConfig
	logger   *jsonlog.Logger
	Services Services
}

func main() {
	cfg := loadConfig()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
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
			Recipes:  service.NewRecipeService(dbProvider),
			Users:    service2.NewUserService(dbProvider),
			Comments: service2.NewCommentService(dbProvider),
		},
	}

	log.Printf("Starting %s API server on port %d", app.config.env, app.config.port)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ErrorLog:     log.New(logger, "", 0),
		IdleTimeout:  cfg.http.idleTimeout,
		ReadTimeout:  cfg.http.readTimeout,
		WriteTimeout: cfg.http.writeTimeout,
	}

	err := server.ListenAndServe()
	logger.PrintFatal(err, nil)
}
