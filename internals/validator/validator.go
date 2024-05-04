package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func ValidateFields(json string, data any) {
	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(data); err != nil {
		log.Fatalf("validator: Failed struct validation | Message: \"%v\" | Provided JSON: \"%v\"", err, json)
	}
}
