package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Env  string
}

func MustLoad() Config {

	if err := godotenv.Load(); err != nil {
		panic("dontenv file not loaded")
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("port is required")
	}

	env := os.Getenv("ENV")
	if env == "" {
		panic("working enviroment is required")
	}

	return Config{
		Port: port,
		Env:  env,
	}
}
