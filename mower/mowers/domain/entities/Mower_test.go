package entities

import (
	"example.kata.local/mower/mowers/domain/exceptions"
	"example.kata.local/mower/mowers/domain/valueobjects"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var genericMower Mower
var genericSurface valueobjects.Surface

func setupTest() {
	mowerId, _ := valueobjects.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobjects.BuildXPosition(5)
	yPosition, _ := valueobjects.BuildYPosition(5)
	orientation, _ := valueobjects.BuildMowerOrientation("N")

	xSize, _ := valueobjects.BuildSurfaceXSize(5)
	ySize, _ := valueobjects.BuildSurfaceYSize(5)
	genericSurface, _ = valueobjects.BuildSurface(xSize, ySize)

	mowerPosition, _ := valueobjects.BuildMowerPosition(xPosition, yPosition, orientation)

	genericMower, _ = BuildMower(mowerId, mowerPosition)
}

func TestMowerShouldBeBuild(t *testing.T) {
	mowerId, _ := valueobjects.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobjects.BuildXPosition(15)
	yPosition, _ := valueobjects.BuildYPosition(8)
	orientation, _ := valueobjects.BuildMowerOrientation("N")

	mowerPosition, _ := valueobjects.BuildMowerPosition(xPosition, yPosition, orientation)

	mower, _ := BuildMower(mowerId, mowerPosition)

	if reflect.TypeOf(mower).String() != "entities.Mower" {
		t.Fatal(reflect.TypeOf(mower))
	}

	assert.Equal(t, mowerPosition, mower.Position())
	assert.Equal(t, mowerId, mower.id())
}

func TestMowerShouldBeAbleToTurnLeft(t *testing.T) {
	setupTest()

	movement, _ := valueobjects.BuildMowerMovement("L")
	expectedOrientation, _ := valueobjects.BuildMowerOrientation("W")

	_ = genericMower.Move(movement, genericSurface)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToTurnRight(t *testing.T) {
	setupTest()

	movement, _ := valueobjects.BuildMowerMovement("R")
	expectedOrientation, _ := valueobjects.BuildMowerOrientation("E")

	_ = genericMower.Move(movement, genericSurface)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToMoveForward(t *testing.T) {
	setupTest()

	movement, _ := valueobjects.BuildMowerMovement("F")
	expectedYPosition, _ := valueobjects.BuildYPosition(6)
	expectedXPosition, _ := valueobjects.BuildXPosition(5)

	_ = genericMower.Move(movement, genericSurface)
	position := genericMower.Position()

	assert.Equal(t, expectedYPosition, position.YPosition())
	assert.Equal(t, expectedXPosition, position.XPosition())
}

func TestMowerShouldFailIfGoesOutOfBounds(t *testing.T) {
	setupTest()

	uuidAsString := uuid.NewString()

	var tests = []struct {
		orientation string
		xPosition   int
		yPosition   int
	}{
		{"N", 5, 5},
		{"S", 5, 0},
		{"E", 5, 5},
		{"W", 0, 5},
	}

	for _, params := range tests {
		testName := fmt.Sprintf("Orientation-%s,xPosition-%d,yPosition-%d", params.orientation, params.xPosition, params.yPosition)

		t.Run(testName, func(t *testing.T) {
			orientation, _ := valueobjects.BuildMowerOrientation(params.orientation)
			xPosition, _ := valueobjects.BuildXPosition(params.xPosition)
			yPosition, _ := valueobjects.BuildYPosition(params.yPosition)

			position, _ := valueobjects.BuildMowerPosition(xPosition, yPosition, orientation)
			mowerId, _ := valueobjects.BuildMowerId(uuidAsString)

			mower, _ := BuildMower(mowerId, position)
			movement, _ := valueobjects.BuildMowerMovement("F")

			err := mower.Move(movement, genericSurface)

			assert.Error(t, err)
			assert.Equal(t, exceptions.BuildMowerOutOfBoundsError(), err)
		})
	}
}
