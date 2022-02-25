package entities

import (
	valueobjects2 "example.kata.local/mower/mowers/domain/valueobjects"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var genericMower Mower

func setupTest() {
	mowerId, _ := valueobjects2.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobjects2.BuildXPosition(5)
	yPosition, _ := valueobjects2.BuildYPosition(5)
	orientation, _ := valueobjects2.BuildMowerOrientation("N")

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

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToTurnRight(t *testing.T) {
	setupTest()

	movement, _ := valueobjects2.BuildMowerMovement("R")
	expectedOrientation, _ := valueobjects2.BuildMowerOrientation("E")

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedOrientation, position.Orientation())
}

func TestMowerShouldBeAbleToMoveForward(t *testing.T) {
	setupTest()

	movement, _ := valueobjects2.BuildMowerMovement("F")
	expectedYPosition, _ := valueobjects2.BuildYPosition(6)
	expectedXPosition, _ := valueobjects2.BuildXPosition(5)

	genericMower.Move(movement)
	position := genericMower.Position()

	assert.Equal(t, expectedYPosition, position.YPosition())
	assert.Equal(t, expectedXPosition, position.XPosition())
}

func TestMowerShouldFailIfGoesOutOfBounds(t *testing.T) {
	assert.True(t, false)
}
