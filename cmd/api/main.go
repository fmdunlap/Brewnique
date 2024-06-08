package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.1"

type application struct {
	config applicationConfig
	logger *log.Logger
}

func main() {
	cfg := loadConfig()

	logger := log.New(os.Stdout, "", log.LstdFlags)

	app := &application{
		config: cfg,
		logger: logger,
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
