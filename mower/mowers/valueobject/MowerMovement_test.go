package valueobject

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMowerMovementShouldBeBuild(t *testing.T) {
	mowerMovement, _ := BuildMowerMovement("L")

	if reflect.TypeOf(mowerMovement).String() != "valueobject.MowerMovement" {
		t.Fatal(reflect.TypeOf(mowerMovement))
	}
	assert.Equal(t, "L", mowerMovement.Value())
}

func TestShouldAcceptValidValues(t *testing.T) {

	var tests = []struct {
		movement string
	}{
		{"L"},
		{"R"},
		{"F"},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("%s", params.movement)

		t.Run(testName, func(t *testing.T) {
			var _, err = BuildMowerMovement(params.movement)

			assert.Nil(t, err)
		})
	}
}

func TestShouldReturnErrorForInvalidValues(t *testing.T) {
	mowerMovement, err := BuildMowerMovement("U")

	// TODO: Has to return specific error
	assert.Equal(t, MowerMovement{}, mowerMovement)
	assert.Error(t, err)
}
