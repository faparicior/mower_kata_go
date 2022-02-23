package valueobject

import (
	"example.kata.local/mower/mowers/exceptions"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
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

func TestShouldReturnErrorForInvalidMovement(t *testing.T) {
	mowerMovement, err := BuildMowerMovement("U")
	expectedError := exceptions.BuildInvalidMowerMovement()

	assert.Equal(t, MowerMovement{}, mowerMovement)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}

func TestShouldEvalIfMovementIsForward(t *testing.T) {

	var tests = []struct {
		movement      string
		expectedValue bool
	}{
		{"L", false},
		{"R", false},
		{"F", true},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("%s expects %s", params.movement, strconv.FormatBool(params.expectedValue))

		t.Run(testName, func(t *testing.T) {
			var movement, _ = BuildMowerMovement(params.movement)

			assert.Equal(t, params.expectedValue, movement.isForward())
		})
	}
}

func TestShouldEvalIfMovementIsClockwise(t *testing.T) {

	var tests = []struct {
		movement      string
		expectedValue bool
	}{
		{"L", false},
		{"R", true},
		{"F", false},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("%s expects %s", params.movement, strconv.FormatBool(params.expectedValue))

		t.Run(testName, func(t *testing.T) {
			var movement, _ = BuildMowerMovement(params.movement)

			assert.Equal(t, params.expectedValue, movement.isClockWise())
		})
	}
}

func TestShouldEvalIfMovementIsCounterClockwise(t *testing.T) {

	var tests = []struct {
		movement      string
		expectedValue bool
	}{
		{"L", true},
		{"R", false},
		{"F", false},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("%s expects %s", params.movement, strconv.FormatBool(params.expectedValue))

		t.Run(testName, func(t *testing.T) {
			var movement, _ = BuildMowerMovement(params.movement)

			assert.Equal(t, params.expectedValue, movement.isCounterClockWise())
		})
	}
}
