package exceptions

import (
	"encoding/json"
	"example.kata.local/mower/internal/shared/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMowerOutOfBoundsShouldBeBuild(t *testing.T) {
	invalidMowerMovement := BuildMowerOutOfBoundsError()

	expectedError, _ := json.Marshal(exceptions.Error{
		Description:  "mower out of bounds",
		ErrorType:    "domain",
		ErrorSubtype: "MowerOutOfBounds",
	})

	assert.Equal(t, string(expectedError), invalidMowerMovement.Error())
}
