package validator

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/josephelias94/tweet-deleter/internals/constants"
)

func ValidateFields(json string, data any) error {
	if err := validator.New(validator.WithRequiredStructEnabled()).Struct(data); err != nil {
		return errors.New(constants.ERROR_VALIDATOR_STRUCT + "Provided JSON: " + json + " | ErrorMessage: " + err.Error())
	}

	return nil
}
