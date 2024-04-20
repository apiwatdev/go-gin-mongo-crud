package bootstraps

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	MONGO_DATABASE       string
	MONGO_CONNECTION_STR string
}

var env *Env

func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	MONGO_DATABASE := os.Getenv("MONGO_DATABASE")
	MONGO_CONNECTION_STR := os.Getenv("MONGO_CONNECTION_STR")
	env = &Env{
		MONGO_DATABASE:       MONGO_DATABASE,
		MONGO_CONNECTION_STR: MONGO_CONNECTION_STR,
	}
}

func GetEnv() *Env {
	return env
}
