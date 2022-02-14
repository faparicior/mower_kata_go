package valueobject

import (
	"example.kata.local/mower/mowers/exceptions"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMowerOrientationShouldBeBuild(t *testing.T) {
	const orientation = "N"

	var mowerOrientation, err = BuildMowerOrientation(orientation)

	if reflect.TypeOf(mowerOrientation).String() != "valueobject.MowerOrientation" {
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

func TestShouldFailForInvalidOrientation(t *testing.T) {
	const invalidOrientation = "J"

	mowerOrientation, err := BuildMowerOrientation(invalidOrientation)

	expectedError := exceptions.BuildInvalidMowerOrientation()

	assert.Equal(t, MowerOrientation{}, mowerOrientation)
	assert.NotNil(t, err)
	assert.Equal(t, expectedError, err)
}
