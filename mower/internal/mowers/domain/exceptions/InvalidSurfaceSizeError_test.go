package exceptions

import (
	"encoding/json"
	"example.kata.local/mower/internal/shared/domain/exceptions"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvalidSurfaceSizeShouldBeBuild(t *testing.T) {
	invalidMowerPosition := BuildInvalidSurfaceSize()

	expectedError, _ := json.Marshal(exceptions.Error{
		Description:  "invalid surface size",
		ErrorType:    "domain",
		ErrorSubtype: "InvalidSurfaceSize",
	})

	assert.Equal(t, string(expectedError), invalidMowerPosition.Error())
}
