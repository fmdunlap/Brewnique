package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type applicationConfig struct {
	port     int
	env      string
	http     httpConfig
	database databaseConfig
}

func (c *applicationConfig) toLogMap() *map[string]string {
	return &map[string]string{
		"port": strconv.Itoa(c.port),
		"env":  c.env,
	}
}

type httpConfig struct {
	idleTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	maxBodySize  int64
}

func (h *httpConfig) String() string {
	return fmt.Sprintf("idleTimeout: %s, readTimeout: %s, writeTimeout: %s, maxBodySize: %d", h.idleTimeout, h.readTimeout, h.writeTimeout, h.maxBodySize)
}

func (h *httpConfig) toLogMap() *map[string]string {
	return &map[string]string{
		"idle_timeout":  h.idleTimeout.String(),
		"read_timeout":  h.readTimeout.String(),
		"write_timeout": h.writeTimeout.String(),
		"max_body_size": fmt.Sprintf("%d", h.maxBodySize),
	}
}

type databaseConfig struct {
	driver          string
	dsn             string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
}

func (d *databaseConfig) toLogMap() *map[string]string {
	return &map[string]string{
		"driver":         d.driver,
		"dsn":            d.dsn,
		"max_open_conns": strconv.Itoa(d.maxOpenConns),
		"max_idle_conns": strconv.Itoa(d.maxIdleConns),
		"conn_max_life":  d.connMaxLifetime.String(),
	}
}

func (d *databaseConfig) String() string {
	return fmt.Sprintf("driver: %s, dsn: %s, maxOpenConns: %d, maxIdleConns: %d, connMaxLifetime: %s", d.driver, d.dsn, d.maxOpenConns, d.maxIdleConns, d.connMaxLifetime)
}

func loadConfig() applicationConfig {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.env")
	viper.SetEnvPrefix("brew_api")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetDefault("port", 8080)
	viper.SetDefault("env", "dev")
	viper.SetDefault("http.idle_timeout", time.Minute)
	viper.SetDefault("http.read_timeout", time.Minute)
	viper.SetDefault("http.write_timeout", time.Minute)
	viper.SetDefault("http.max_body_size", 4*1024*1024)

	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.dsn", "postgres://postgres:postgres@localhost:5432/brewnique")
	viper.SetDefault("database.max_open_conns", 10)
	viper.SetDefault("database.max_idle_conns", 10)
	viper.SetDefault("database.conn_max_lifetime", time.Minute*15)

	viper.AutomaticEnv()

	return applicationConfig{
		port: viper.GetInt("port"),
		env:  viper.GetString("env"),
		http: httpConfig{
			idleTimeout:  viper.GetDuration("http.idle_timeout"),
			readTimeout:  viper.GetDuration("http.read_timeout"),
			writeTimeout: viper.GetDuration("http.write_timeout"),
			maxBodySize:  viper.GetInt64("http.max_body_size"),
		},
		database: databaseConfig{
			driver:          viper.GetString("database.driver"),
			dsn:             viper.GetString("database.dsn"),
			maxIdleConns:    viper.GetInt("database.max_idle_conns"),
			maxOpenConns:    viper.GetInt("database.max_open_conns"),
			connMaxLifetime: viper.GetDuration("database.conn_max_lifetime"),
		},
	}
}
