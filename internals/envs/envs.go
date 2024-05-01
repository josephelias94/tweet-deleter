package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("envs: Error loading .env file. Message: \"%v\"", err)
	}
}

func getValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("envs: Error loading \"%v\" variable. Check if .env file is loaded or the environment variable is not empty.", key)
	}

	return value
}

func GetBearerToken() string {
	return getValue("BEARER_TOKEN")
}
