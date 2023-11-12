package config

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type config struct {
	// Default data
	ServerPort string `env:"SERVER_PORT" envDefault:":3000"`
	PgUser     string `env:"PG_USER"`
	PgPassword string `env:"PG_PASSWORD"`
	PgDBName   string `env:"PG_DBNAME"`
	PgHost     string `env:"PG_HOST"`
	JWTSecret  string `env:"JWT_SECRET"`
	BucketName string `env:"BUCKET_NAME"`
}

func new() config {
	_ = godotenv.Load()
	settings := config{}
	if err := env.Parse(&settings); err != nil {
		panic(
			fmt.Sprintf(
				"Error in environment variables:\n\n%s", err.Error(),
			),
		)
	}
	return settings
}

var Settings = new()
