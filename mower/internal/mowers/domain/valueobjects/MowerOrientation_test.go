package valueobjects

import (
	"example.kata.local/mower/internal/mowers/domain/exceptions"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
)

func TestMowerOrientationShouldBeBuild(t *testing.T) {
	const orientation = "N"

	var mowerOrientation, err = BuildMowerOrientation(orientation)

	if reflect.TypeOf(mowerOrientation).String() != "valueobjects.MowerOrientation" {
		t.Fatal(reflect.TypeOf(mowerOrientation))
	}

	assert.Nil(t, err)
}

func TestMowerOrientationShouldBeBuildWithValidParameters(t *testing.T) {

	var tests = []struct {
		orientation string
	}{
		{"N"},
		{"S"},
		{"E"},
		{"W"},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("%s", params.orientation)

		t.Run(testName, func(t *testing.T) {
			var _, err = BuildMowerOrientation(params.orientation)

			assert.Nil(t, err)
		})
	}
}

func TestMowerOrientationShouldFailForInvalidOrientation(t *testing.T) {
	const invalidOrientation = "J"

	mowerOrientation, err := BuildMowerOrientation(invalidOrientation)

	expectedError := exceptions.BuildInvalidMowerOrientation()

	assert.Equal(t, MowerOrientation{}, mowerOrientation)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}

func TestMowerOrientationShouldApplyMovementOrientations(t *testing.T) {
	var tests = []struct {
		orientation         string
		movement            string
		expectedOrientation string
	}{
		{"N", "F", "N"},
		{"N", "R", "E"},
		{"E", "R", "S"},
		{"S", "R", "W"},
		{"W", "R", "N"},
		{"N", "L", "W"},
		{"W", "L", "S"},
		{"S", "L", "E"},
		{"E", "L", "N"},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("Orientation-%s,Movement-%s,Expected-%s", params.orientation, params.movement, params.expectedOrientation)

		t.Run(testName, func(t *testing.T) {
			orientation, _ := BuildMowerOrientation(params.orientation)

			movement, _ := BuildMowerMovement(params.movement)
			orientation = orientation.ChangeOrientation(movement)

			assert.Equal(t, params.expectedOrientation, orientation.value)
		})
	}
}

func TestMowerOrientationShouldEvalIfAffectsYAxis(t *testing.T) {
	var tests = []struct {
		orientation         string
		expectedAffectYAxis bool
	}{
		{"N", true},
		{"S", true},
		{"E", false},
		{"W", false},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("Orientation-%s,Expected-%s", params.orientation, strconv.FormatBool(params.expectedAffectYAxis))

		t.Run(testName, func(t *testing.T) {
			orientation, _ := BuildMowerOrientation(params.orientation)

			assert.Equal(t, params.expectedAffectYAxis, orientation.AffectsYAxis())
		})
	}
}

func TestMowerOrientationShouldEvalIfAffectsXAxis(t *testing.T) {
	var tests = []struct {
		orientation         string
		expectedAffectYAxis bool
	}{
		{"N", false},
		{"S", false},
		{"E", true},
		{"W", true},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("Orientation-%s,Expected-%s", params.orientation, strconv.FormatBool(params.expectedAffectYAxis))

		t.Run(testName, func(t *testing.T) {
			orientation, _ := BuildMowerOrientation(params.orientation)

			assert.Equal(t, params.expectedAffectYAxis, orientation.AffectsXAxis())
		})
	}
}

func TestMowerOrientationShouldEvalDirectionStep(t *testing.T) {
	var tests = []struct {
		orientation  string
		expectedStep int
	}{
		{"N", 1},
		{"E", 1},
		{"S", -1},
		{"W", -1},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("Orientation-%s,Expected-%d", params.orientation, params.expectedStep)

		t.Run(testName, func(t *testing.T) {
			orientation, _ := BuildMowerOrientation(params.orientation)

			assert.Equal(t, params.expectedStep, orientation.StepMovement())
		})
	}
}
