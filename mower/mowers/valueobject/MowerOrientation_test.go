package valueobject

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const orientation = "N"
const invalidOrientation = "J"

func TestMowerOrientationShouldBeBuild(t *testing.T) {

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
	mowerOrientation, err := BuildMowerOrientation(invalidOrientation)

	// TODO: Has to return specific error
	assert.Equal(t, MowerOrientation{}, mowerOrientation)
	assert.NotNil(t, err)
}
