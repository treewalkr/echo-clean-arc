package config

import (
	"os"
	"sync"

	godotenv "github.com/joho/Godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DB  DBConfig
	App AppConfig
}

type DBConfig struct {
	DSN string
}

type AppConfig struct {
	Port string
}

var (
	cfg  *Config
	once sync.Once
)

func Load() (*Config, error) {
	godotenv.Load()
	env := os.Getenv("ENV")

	if env == "development" {
		viper.SetConfigFile(".env." + env)
		viper.ReadInConfig()

		// Set default values
		viper.SetDefault("DB_DSN", "example.db")
		viper.SetDefault("APP_PORT", "3000")

		once.Do(func() {
			cfg = &Config{
				DB: DBConfig{
					DSN: viper.GetString("DB_DSN"),
				},
				App: AppConfig{
					Port: os.Getenv("APP_PORT"),
				},
			}
		})
		return cfg, nil
	}

	// TODO: Implement other environments
	return cfg, nil
}
