package valueobjects

import (
	"example.kata.local/mower/internal/mowers/domain/exceptions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXPositionShouldBeBuild(t *testing.T) {
	const mowerPosition = 15

	xPosition, err := BuildXPosition(mowerPosition)

	if reflect.TypeOf(xPosition).String() != "valueobjects.XPosition" {
		t.Fatal(reflect.TypeOf(xPosition))
	}

	assert.Equal(t, mowerPosition, xPosition.Value())
	//if xPosition.Value() != 15 {
	//	t.Fatal("error")
	//}

	assert.Nil(t, err)
}

func TestXPositionShouldThrownExceptionForNegativeValues(t *testing.T) {
	const invalidMowerPosition = -1

	xPosition, err := BuildXPosition(invalidMowerPosition)
	expectedError := exceptions.BuildInvalidMowerPosition()

	assert.Equal(t, XPosition{}, xPosition)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}

func TestXPositionShouldSumAStepForward(t *testing.T) {
	const position = 15
	const step = 1

	xPosition, _ := BuildXPosition(position)
	xPosition, _ = xPosition.MoveForward(step)

	assert.Equal(t, position+step, xPosition.Value())
}

func TestXPositionShouldFailForNegativeValuesOnStepForward(t *testing.T) {
	const position = 0
	const step = -1
	expectedError := exceptions.BuildInvalidMowerPosition()

	xPosition, _ := BuildXPosition(position)
	xPosition, err := xPosition.MoveForward(step)

	assert.Equal(t, XPosition{}, xPosition)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}
