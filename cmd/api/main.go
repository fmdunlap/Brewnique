package main

import (
	"flag"
	"fmt"
	"github.com/alecthomas/units"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "0.0.1"

const DefaultPort = 8080
const DefaultEnv = "dev"
const DefaultIdleTimeout = 60 * time.Second
const DefaultReadTimeout = 60 * time.Second
const DefaultWriteTimeout = 60 * time.Second
const DefaultMaxBodySize = 4 * units.Mebibyte

type application struct {
	config applicationConfig
	logger *log.Logger
}

func main() {
	cfg := applicationConfig{
		port: DefaultPort,
		env:  DefaultEnv,
		http: httpConfig{
			idleTimeout:  DefaultIdleTimeout,
			readTimeout:  DefaultReadTimeout,
			writeTimeout: DefaultWriteTimeout,
			maxBodySize:  int64(DefaultMaxBodySize),
		},
	}

	flag.IntVar(&cfg.port, "port", DefaultPort, "Port to listen on")
	flag.StringVar(&cfg.env, "env", DefaultEnv, "Environment to use (dev, prod)")
	flag.Parse()

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
