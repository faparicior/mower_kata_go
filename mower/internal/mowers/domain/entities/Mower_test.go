package entities

import (
	"example.kata.local/mower/internal/mowers/domain/exceptions"
	valueobjects2 "example.kata.local/mower/internal/mowers/domain/valueobjects"
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var genericMower Mower
var genericSurface valueobjects2.Surface

func setupTest() {
	mowerId, _ := valueobjects2.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobjects2.BuildXPosition(5)
	yPosition, _ := valueobjects2.BuildYPosition(5)
	orientation, _ := valueobjects2.BuildMowerOrientation("N")

	xSize, _ := valueobjects2.BuildSurfaceXSize(5)
	ySize, _ := valueobjects2.BuildSurfaceYSize(5)
	genericSurface, _ = valueobjects2.BuildSurface(xSize, ySize)

	mowerPosition, _ := valueobjects2.BuildMowerPosition(xPosition, yPosition, orientation)

	genericMower, _ = BuildMower(mowerId, mowerPosition)
}

func TestMowerShouldBeBuild(t *testing.T) {
	mowerId, _ := valueobjects2.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobjects2.BuildXPosition(15)
	yPosition, _ := valueobjects2.BuildYPosition(8)
	orientation, _ := valueobjects2.BuildMowerOrientation("N")

	mowerPosition, _ := valueobjects2.BuildMowerPosition(xPosition, yPosition, orientation)

	mower, _ := BuildMower(mowerId, mowerPosition)

	if reflect.TypeOf(mower).String() != "entities.Mower" {
		t.Fatal(reflect.TypeOf(mower))
	}

	assert.Equal(t, mowerPosition, mower.Position())
	assert.Equal(t, mowerId, mower.id())
}

func TestMowerShouldBeAbleToTurnLeft(t *testing.T) {
	setupTest()

	movement, _ := valueobjects2.BuildMowerMovement("L")
	expectedOrientation, _ := valueobjects2.BuildMowerOrientation("W")

	_ = genericMower.Move(movement, genericSurface)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToTurnRight(t *testing.T) {
	setupTest()

	movement, _ := valueobjects2.BuildMowerMovement("R")
	expectedOrientation, _ := valueobjects2.BuildMowerOrientation("E")

	_ = genericMower.Move(movement, genericSurface)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToMoveForward(t *testing.T) {
	setupTest()

	movement, _ := valueobjects2.BuildMowerMovement("F")
	expectedYPosition, _ := valueobjects2.BuildYPosition(6)
	expectedXPosition, _ := valueobjects2.BuildXPosition(5)

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
			orientation, _ := valueobjects2.BuildMowerOrientation(params.orientation)
			xPosition, _ := valueobjects2.BuildXPosition(params.xPosition)
			yPosition, _ := valueobjects2.BuildYPosition(params.yPosition)

			position, _ := valueobjects2.BuildMowerPosition(xPosition, yPosition, orientation)
			mowerId, _ := valueobjects2.BuildMowerId(uuidAsString)

			mower, _ := BuildMower(mowerId, position)
			movement, _ := valueobjects2.BuildMowerMovement("F")

			err := mower.Move(movement, genericSurface)

			assert.Error(t, err)
			assert.Equal(t, exceptions.BuildMowerOutOfBoundsError(), err)
		})
	}
}
