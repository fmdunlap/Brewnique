package main

import (
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

type applicationConfig struct {
	port     int
	env      string
	http     httpConfig
	database databaseConfig
}

type httpConfig struct {
	idleTimeout  time.Duration
	readTimeout  time.Duration
	writeTimeout time.Duration
	maxBodySize  int64
}

type databaseConfig struct {
	driver          string
	dsn             string
	maxOpenConns    int
	maxIdleConns    int
	connMaxLifetime time.Duration
}

func loadConfig() applicationConfig {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.env")
	viper.SetEnvPrefix("brewnique")
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

func initConfigFromEnv() applicationConfig {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.SetEnvPrefix("brewnique")

	log.Printf(viper.ConfigFileUsed())

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file:", err)
	}

	config := applicationConfig{
		port: viper.GetInt("port"),
		env:  viper.GetString("env"),
		http: httpConfig{
			idleTimeout:  viper.GetDuration("http.idleTimeout"),
			readTimeout:  viper.GetDuration("http.readTimeout"),
			writeTimeout: viper.GetDuration("http.writeTimeout"),
			maxBodySize:  viper.GetInt64("http.maxBodySize"),
		},
	}

	return config
}
