package entities

import (
	"example.kata.local/mower/mowers/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var genericMower Mower

func setupTest() {
	mowerId, _ := valueobjects.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobjects.BuildXPosition(5)
	yPosition, _ := valueobjects.BuildYPosition(5)
	orientation, _ := valueobjects.BuildMowerOrientation("N")

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

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToTurnRight(t *testing.T) {
	setupTest()

	movement, _ := valueobjects.BuildMowerMovement("R")
	expectedOrientation, _ := valueobjects.BuildMowerOrientation("E")

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToMoveForward(t *testing.T) {
	setupTest()

	movement, _ := valueobjects.BuildMowerMovement("F")
	expectedYPosition, _ := valueobjects.BuildYPosition(6)
	expectedXPosition, _ := valueobjects.BuildXPosition(5)

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedYPosition, position.YPosition())
	assert.Equal(t, expectedXPosition, position.XPosition())
}

func TestMowerShouldFailIfGoesOutOfBounds(t *testing.T) {
	assert.True(t, false)
}
