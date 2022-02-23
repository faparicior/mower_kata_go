package valueobject

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMowerPositionShouldBeBuild(t *testing.T) {
	xPosition, _ := BuildXPosition(15)
	yPosition, _ := BuildYPosition(8)
	orientation, _ := BuildMowerOrientation("N")

	mowerPosition, _ := BuildMowerPosition(xPosition, yPosition, orientation)

	if reflect.TypeOf(mowerPosition).String() != "valueobject.MowerPosition" {
		t.Fatal(reflect.TypeOf(xPosition))
	}

	assert.Equal(t, xPosition, mowerPosition.XPosition())
	assert.Equal(t, yPosition, mowerPosition.YPosition())
	assert.Equal(t, orientation, mowerPosition.Orientation())
}

func TestShouldChangeOrientationWithRightMovement(t *testing.T) {
	xPosition, _ := BuildXPosition(15)
	yPosition, _ := BuildYPosition(8)
	orientation, _ := BuildMowerOrientation("N")

	mowerPosition, _ := BuildMowerPosition(xPosition, yPosition, orientation)
	mowerMovement, _ := BuildMowerMovement("R")

	newMowerPosition, _ := mowerPosition.Move(mowerMovement)
	expectedMowerOrientation, _ := BuildMowerOrientation("E")

	assert.Equal(t, expectedMowerOrientation, newMowerPosition.Orientation())
}

func TestShouldChangeOrientationWithLeftMovement(t *testing.T) {
	xPosition, _ := BuildXPosition(15)
	yPosition, _ := BuildYPosition(8)
	orientation, _ := BuildMowerOrientation("N")

	mowerPosition, _ := BuildMowerPosition(xPosition, yPosition, orientation)
	mowerMovement, _ := BuildMowerMovement("L")

	newMowerPosition, _ := mowerPosition.Move(mowerMovement)
	expectedMowerOrientation, _ := BuildMowerOrientation("W")

	assert.Equal(t, expectedMowerOrientation, newMowerPosition.Orientation())
}

func TestShouldBeAbleToChangePosition(t *testing.T) {

	var tests = []struct {
		orientation       string
		expectedXPosition int
		expectedYPosition int
	}{
		{"N", 5, 6},
		{"E", 6, 5},
		{"S", 5, 4},
		{"W", 4, 5},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("Orientation %s expects X-%d, Y-%d", params.orientation, params.expectedXPosition, params.expectedYPosition)

		xPosition, _ := BuildXPosition(5)
		yPosition, _ := BuildYPosition(5)
		orientation, _ := BuildMowerOrientation(params.orientation)

		mowerPosition, _ := BuildMowerPosition(xPosition, yPosition, orientation)

		t.Run(testName, func(t *testing.T) {
			var movement, _ = BuildMowerMovement("F")

			newMowerPosition, _ := mowerPosition.Move(movement)
			assert.Equal(t, params.expectedXPosition, newMowerPosition.XPosition().Value())
			assert.Equal(t, params.expectedYPosition, newMowerPosition.YPosition().Value())
		})
	}
}
