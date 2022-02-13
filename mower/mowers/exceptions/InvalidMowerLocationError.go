package exceptions

import (
	"encoding/json"
	"errors"
	"example.kata.local/mower/shared/domain/exceptions"
)

func BuildInvalidMowerLocation() error {
	customError, err := json.Marshal(exceptions.Error{
		Description:  "negative position",
		ErrorType:    "Domain",
		ErrorSubtype: "InvalidMowerLocation",
	})

	if err != nil {
		panic(err)
	}

	return errors.New(string(customError))
}
