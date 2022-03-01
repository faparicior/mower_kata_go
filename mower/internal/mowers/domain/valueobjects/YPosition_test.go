package valueobjects

import (
	"example.kata.local/mower/internal/mowers/domain/exceptions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYPositionShouldBeBuild(t *testing.T) {
	const mowerPosition = 15

	yPosition, err := BuildYPosition(mowerPosition)

	if reflect.TypeOf(yPosition).String() != "valueobjects.YPosition" {
		t.Fatal(reflect.TypeOf(yPosition))
	}

	assert.Equal(t, mowerPosition, yPosition.Value())
	//if yPosition.Value() != 15 {
	//	t.Fatal("error")
	//}

	assert.Nil(t, err)
}

func TestYPositionShouldThrownExceptionForNegativeValues(t *testing.T) {
	const invalidMowerPosition = -1

	yPosition, err := BuildYPosition(invalidMowerPosition)
	expectedError := exceptions.BuildInvalidMowerPosition()

	assert.Equal(t, YPosition{}, yPosition)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}

func TestYPositionShouldSumAStepForward(t *testing.T) {
	const position = 15
	const step = 1

	yPosition, _ := BuildYPosition(position)
	yPosition, _ = yPosition.MoveForward(step)

	assert.Equal(t, position+step, yPosition.Value())
}

func TestYPositionShouldFailForNegativeValuesOnStepForward(t *testing.T) {
	const position = 0
	const step = -1
	expectedError := exceptions.BuildInvalidMowerPosition()

	yPosition, _ := BuildYPosition(position)
	yPosition, err := yPosition.MoveForward(step)

	assert.Equal(t, YPosition{}, yPosition)
	assert.Error(t, err)
	assert.EqualValues(t, expectedError, err)
}
