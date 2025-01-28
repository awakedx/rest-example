package config

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	PgURL            string
	HTTPAddr         string
	ReadTimeOut      time.Duration
	WriteTimeOut     time.Duration
	PgMigrationsPath string
	AccessTokenTTL   time.Duration
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			slog.Error("failed to load .env file", "error", err.Error())
			panic(err)
		}
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("../../config/")
		viper.AddConfigPath("./config")

		err = viper.ReadInConfig()
		if err != nil {
			slog.Error("failed to read .yaml file", "error", err.Error())
			panic(err)
		}

		config = Config{
			PgURL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
				os.Getenv("DB_USER"),
				os.Getenv("DB_PASSWORD"),
				os.Getenv("DB_HOST"),
				os.Getenv("DB_PORT"),
				os.Getenv("DB_NAME")),
			HTTPAddr:         "0.0.0.0:" + viper.GetString("server.port"),
			ReadTimeOut:      viper.GetDuration("server.readTimeOut"),
			WriteTimeOut:     viper.GetDuration("server.writeTimeOut"),
			PgMigrationsPath: viper.GetString("database.migrationPath"),
			AccessTokenTTL:   viper.GetDuration("auth.AccessTokenTTL"),
		}
		slog.Info("Config was successfully read")
	})
	return &config
}
