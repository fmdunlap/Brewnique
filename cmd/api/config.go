package main

import "time"

type applicationConfig struct {
	port int
	env  string
	http httpConfig
}

type httpConfig struct {
	idleTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	maxBodySize  int64
}

type databaseConfig struct {
	driver string
	dsn    string
}

func initConfigFromEnv() applicationConfig {
	return applicationConfig{
		port: 8080,
		env:  "dev",
		http: httpConfig{
			idleTimeout:  time.Minute,
			readTimeout:  time.Minute,
			writeTimeout: time.Minute,
			maxBodySize:  4 * 1024 * 1024,
		},
	}
}
