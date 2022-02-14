package valueobject

import (
	"example.kata.local/mower/mowers/exceptions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mowerPosition = 15
const invalidMowerPosition = -1

const position = 15
const step = 1

func TestXPositionShouldBeBuild(t *testing.T) {
	var xPosition, err = BuildXPosition(mowerPosition)

	if reflect.TypeOf(xPosition).String() != "valueobject.XPosition" {
		t.Fatal(reflect.TypeOf(xPosition))
	}

	assert.Equal(t, mowerPosition, xPosition.Value())
	//if xPosition.Value() != 15 {
	//	t.Fatal("error")
	//}

	assert.Nil(t, err)
}

func TestXPositionShouldThrownExceptionForNegativeValues(t *testing.T) {
	var xPosition, err = BuildXPosition(invalidMowerPosition)
	var expectedError = exceptions.BuildInvalidMowerPosition()

	assert.Equal(t, XPosition{}, xPosition)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}

func TestXPositionShouldSumAStepForward(t *testing.T) {
	var xPosition, _ = BuildXPosition(position)

	xPosition, _ = xPosition.MoveForward(step)

	assert.Equal(t, position+step, xPosition.Value())
}
