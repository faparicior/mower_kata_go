package exceptions

import (
	"encoding/json"
	"example.kata.local/mower/shared/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidMowerOrientationShouldBeBuild(t *testing.T) {
	invalidMowerOrientation := BuildInvalidMowerOrientation()

	expectedError, _ := json.Marshal(exceptions.Error{
		Description:  "negative position",
		ErrorType:    "domain",
		ErrorSubtype: "InvalidMowerOrientation",
	})

	assert.Equal(t, string(expectedError), invalidMowerOrientation.Error())
}
