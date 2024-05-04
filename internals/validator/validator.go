package validator

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/josephelias94/tweet-deleter/internals/constants"
)

func ValidateFields(json string, data any) {
	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(data); err != nil {
		log.Fatalf("%v Message: \"%v\" | Provided JSON: \"%v\"",
			constants.ERROR_VALIDATOR_STRUCT, err, json)
	}
}
