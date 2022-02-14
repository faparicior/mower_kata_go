package exceptions

import (
	"encoding/json"
	"errors"
	"example.kata.local/mower/shared/domain/exceptions"
)

func BuildInvalidMowerMovement() error {
	const description = "negative position"
	const errorSubtype = "InvalidMowerMovement"

	customError, err := json.Marshal(exceptions.Error{
		Description:  description,
		ErrorType:    exceptions.ErrorType,
		ErrorSubtype: errorSubtype,
	})

	if err != nil {
		panic(err)
	}

	return errors.New(string(customError))
}
