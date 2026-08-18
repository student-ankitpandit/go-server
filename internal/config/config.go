package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
	DatabaseUrl string
}

func MustLoad() Config { //must pattern
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		panic("PORT is required!")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("ENV is required!")
	}

	DBUrl := os.Getenv("DATABASE_URL")
	if DBUrl == "" {
		panic("Database Url is required")
	}

	return Config{
		Port: port,
		Env:  env,
		DatabaseUrl: DBUrl,
	}
}
