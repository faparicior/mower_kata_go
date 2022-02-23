package entities

import (
	"example.kata.local/mower/mowers/valueobject"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var genericMower Mower

func setupTest() {
	mowerId, _ := valueobject.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobject.BuildXPosition(5)
	yPosition, _ := valueobject.BuildYPosition(5)
	orientation, _ := valueobject.BuildMowerOrientation("N")

	mowerPosition, _ := valueobject.BuildMowerPosition(xPosition, yPosition, orientation)

	genericMower, _ = BuildMower(mowerId, mowerPosition)
}

func TestMowerShouldBeBuild(t *testing.T) {
	mowerId, _ := valueobject.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobject.BuildXPosition(15)
	yPosition, _ := valueobject.BuildYPosition(8)
	orientation, _ := valueobject.BuildMowerOrientation("N")

	mowerPosition, _ := valueobject.BuildMowerPosition(xPosition, yPosition, orientation)

	mower, _ := BuildMower(mowerId, mowerPosition)

	if reflect.TypeOf(mower).String() != "entities.Mower" {
		t.Fatal(reflect.TypeOf(mower))
	}

	assert.Equal(t, mowerPosition, mower.Position())
	assert.Equal(t, mowerId, mower.id())
}

func TestMowerShouldBeAbleToTurnLeft(t *testing.T) {
	setupTest()

	movement, _ := valueobject.BuildMowerMovement("L")
	expectedOrientation, _ := valueobject.BuildMowerOrientation("W")

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToTurnRight(t *testing.T) {
	setupTest()

	movement, _ := valueobject.BuildMowerMovement("R")
	expectedOrientation, _ := valueobject.BuildMowerOrientation("E")

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToMoveForward(t *testing.T) {
	setupTest()

	movement, _ := valueobject.BuildMowerMovement("F")
	expectedYPosition, _ := valueobject.BuildYPosition(6)
	expectedXPosition, _ := valueobject.BuildXPosition(5)

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedYPosition, position.YPosition())
	assert.Equal(t, expectedXPosition, position.XPosition())
}

func TestMowerShouldFailIfGoesOutOfBounds(t *testing.T) {
	assert.True(t, false)
}
