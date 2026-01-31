package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type DB struct {
	Host           string `mapstructure:"host" validate:"required"`
	Port           int    `mapstructure:"port" validate:"required,gt=0"`
	User           string `mapstructure:"user" validate:"required"`
	Password       string `mapstructure:"password" validate:"required"`
	Name           string `mapstructure:"name" validate:"required"`
	MaxOpen        int    `mapstructure:"max_open" validate:"required"`
	MaxIdle        int    `mapstructure:"max_idle" validate:"required"`
	MaxConLifetime int    `mapstructure:"max_con_lifetime" validate:"gte=0"`
	SslMode        string `mapstructure:"ssl_mode" validate:"required"`
}

type Server struct {
	Port int `mapstructure:"port" validate:"required,gt=0,lt=65536"`
}

type Config struct {
	DB     *DB     `mapstructure:"db"`
	Server *Server `mapstructure:"server"`
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	appEnv := "develop"
	if key, ok := os.LookupEnv("APP_ENV"); ok {
		appEnv = key
	}

	viper.SetConfigName(appEnv)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	validate := validator.New()
	if err := validate.Struct(&cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return &cfg, nil
}
