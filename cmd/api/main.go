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

type config struct {
	port         int
	env          string
	idleTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	maxBodySize  int64
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	cfg := config{
		port:         DefaultPort,
		env:          DefaultEnv,
		idleTimeout:  DefaultIdleTimeout,
		readTimeout:  DefaultReadTimeout,
		writeTimeout: DefaultWriteTimeout,
		maxBodySize:  int64(DefaultMaxBodySize),
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
		IdleTimeout:  cfg.idleTimeout,
		ReadTimeout:  cfg.readTimeout,
		WriteTimeout: cfg.writeTimeout,
	}

	err := server.ListenAndServe()
	logger.Fatal(err)
}
