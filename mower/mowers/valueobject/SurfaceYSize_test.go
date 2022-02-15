package valueobject

import (
	"example.kata.local/mower/mowers/exceptions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurfaceYSizeShouldBeBuild(t *testing.T) {
	const mowerPosition = 15

	surfaceYSize, err := BuildSurfaceYSize(mowerPosition)

	if reflect.TypeOf(surfaceYSize).String() != "valueobject.SurfaceYSize" {
		t.Fatal(reflect.TypeOf(surfaceYSize))
	}

	assert.Equal(t, mowerPosition, surfaceYSize.Value())
	//if surfaceYSize.Value() != 15 {
	//	t.Fatal("error")
	//}

	assert.Nil(t, err)
}

func TestSurfaceYSizeShouldThrownExceptionForNegativeValues(t *testing.T) {
	const invalidSurfaceSize = -1

	surfaceYSize, err := BuildSurfaceYSize(invalidSurfaceSize)
	expectedError := exceptions.BuildInvalidSurfaceSize()

	assert.Equal(t, SurfaceYSize{}, surfaceYSize)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}
