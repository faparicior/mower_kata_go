package exceptions

import (
	"encoding/json"
	"example.kata.local/mower/shared/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidMowerMovementShouldBeBuild(t *testing.T) {
	invalidMowerMovement := BuildInvalidMowerMovement()

	expectedError, _ := json.Marshal(exceptions.Error{
		Description:  "negative position",
		ErrorType:    "domain",
		ErrorSubtype: "InvalidMowerMovement",
	})

	assert.Equal(t, string(expectedError), invalidMowerMovement.Error())
}
