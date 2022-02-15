package valueobject

import (
	"example.kata.local/mower/mowers/exceptions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurfaceXSizeShouldBeBuild(t *testing.T) {
	const mowerPosition = 15

	surfaceXSize, err := BuildSurfaceXSize(mowerPosition)

	if reflect.TypeOf(surfaceXSize).String() != "valueobject.SurfaceXSize" {
		t.Fatal(reflect.TypeOf(surfaceXSize))
	}

	assert.Equal(t, mowerPosition, surfaceXSize.Value())
	//if surfaceXSize.Value() != 15 {
	//	t.Fatal("error")
	//}

	assert.Nil(t, err)
}

func TestSurfaceXSizeShouldThrownExceptionForNegativeValues(t *testing.T) {
	const invalidSurfaceSize = -1

	surfaceXSize, err := BuildSurfaceXSize(invalidSurfaceSize)
	expectedError := exceptions.BuildInvalidSurfaceSize()

	assert.Equal(t, SurfaceXSize{}, surfaceXSize)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}
