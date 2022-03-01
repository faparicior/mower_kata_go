package exceptions

import (
	"encoding/json"
	"errors"
	"example.kata.local/mower/internal/shared/domain/exceptions"
)

func BuildInvalidSurfaceSize() error {
	const description = "invalid surface size"
	const errorSubtype = "InvalidSurfaceSize"

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
