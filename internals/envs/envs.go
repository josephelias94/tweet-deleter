package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/josephelias94/tweet-deleter/internals/constants"
)

func Load() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("%v Message: \"%v\"", constants.ERROR_ENVS_LOADING_FILE, err)
	}
}

func getValue(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("%v Key: \"%v\"", constants.ERROR_ENVS_LOADING_VARIABLE, key)
	}

	return value
}

func GetBearerToken() string {
	return getValue("BEARER_TOKEN")
}
