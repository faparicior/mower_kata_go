package valueobject

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMowerOrientationShouldBeBuild(t *testing.T) {

	var mowerOrientation, err = BuildMowerOrientation("N")

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

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.orientation)

		t.Run(testName, func(t *testing.T) {
			var _, err = BuildMowerOrientation(tt.orientation)

			assert.Nil(t, err)
		})
	}
}

func TestShouldFailForInvalidOrientation(t *testing.T) {
	var _, err = BuildMowerOrientation("J")

	assert.NotNil(t, err)
}
