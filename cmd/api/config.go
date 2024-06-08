package main

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

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

func loadConfig() applicationConfig {
	viper.SetDefault("port", 8080)
	viper.SetDefault("env", "dev")
	viper.SetDefault("http.idleTimeout", time.Minute)
	viper.SetDefault("http.readTimeout", time.Minute)
	viper.SetDefault("http.writeTimeout", time.Minute)
	viper.SetDefault("http.maxBodySize", 4*1024*1024)

	viper.SetDefault("database.driver", "postgres")
	viper.SetDefault("database.dsn", "postgres://postgres:postgres@localhost:5432/brewnique")

	viper.AddConfigPath(".")
	viper.SetConfigFile("config.env")
	viper.SetEnvPrefix("brewnique")

	viper.AutomaticEnv()

	return applicationConfig{
		port: viper.GetInt("port"),
		env:  viper.GetString("env"),
		http: httpConfig{
			idleTimeout:  viper.GetDuration("http.idleTimeout"),
			readTimeout:  viper.GetDuration("http.readTimeout"),
			writeTimeout: viper.GetDuration("http.writeTimeout"),
			maxBodySize:  viper.GetInt64("http.maxBodySize"),
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
