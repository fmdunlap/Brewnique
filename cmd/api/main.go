package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"brewnique.fdunlap.com/internal/extern/psql"
	"brewnique.fdunlap.com/internal/jsonlog"
	"brewnique.fdunlap.com/internal/service"
)

const version = "0.0.1"

type Services struct {
	Recipes  *service.RecipeService
	Users    *service.UserService
	Comments *service.CommentService
}

type application struct {
	config   applicationConfig
	logger   *jsonlog.Logger
	Services Services
}

func main() {
	cfg := loadConfig()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo, "main")
	logger.PrintInfo("Creating DB provider", cfg.database.toLogMap())
	dbProvider := psql.NewPsqlProvider(psql.PsqlConfig{
		Dsn:             cfg.database.dsn,
		MaxOpenConns:    cfg.database.maxOpenConns,
		MaxIdleConns:    cfg.database.maxIdleConns,
		ConnMaxLifetime: cfg.database.connMaxLifetime,
	})
	logger.PrintInfo("Created DB provider", nil)
	defer dbProvider.Close()

	logger.PrintInfo("Loaded configuration", cfg.toLogMap())
	app := &application{
		config: cfg,
		logger: logger,
		Services: Services{
			Recipes:  service.NewRecipeService(dbProvider),
			Users:    service.NewUserService(dbProvider),
			Comments: service.NewCommentService(dbProvider),
		},
	}
	logger.PrintInfo("Starting API server", &map[string]string{
		"port": strconv.Itoa(app.config.port),
		"env":  app.config.env,
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ErrorLog:     log.New(logger, "", 0),
		IdleTimeout:  cfg.http.idleTimeout,
		ReadTimeout:  cfg.http.readTimeout,
		WriteTimeout: cfg.http.writeTimeout,
	}

	log.Printf("Starting %s API server on port %d", app.config.env, app.config.port)
	err := server.ListenAndServe()
	logger.PrintFatal(err, nil)
}
