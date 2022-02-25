package exceptions

import (
	"encoding/json"
	"example.kata.local/mower/shared/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidMowerPositionShouldBeBuild(t *testing.T) {
	invalidMowerPosition := BuildInvalidMowerPosition()

	expectedError, _ := json.Marshal(exceptions.Error{
		Description:  "negative position",
		ErrorType:    "domain",
		ErrorSubtype: "InvalidMowerPosition",
	})

	assert.Equal(t, string(expectedError), invalidMowerPosition.Error())
}
