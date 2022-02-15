package entities

import (
	"example.kata.local/mower/mowers/valueobject"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMowerShouldBeBuild(t *testing.T) {
	mowerId, _ := valueobject.BuildMowerId(uuid.NewString())
	xPosition, _ := valueobject.BuildXPosition(15)
	yPosition, _ := valueobject.BuildYPosition(8)
	mowerPosition, _ := valueobject.BuildMowerPosition(xPosition, yPosition)

	mower, _ := BuildMower(mowerId, mowerPosition)

	if reflect.TypeOf(mower).String() != "entities.Mower" {
		t.Fatal(reflect.TypeOf(mower))
	}

	assert.Equal(t, mowerPosition, mower.position())
	assert.Equal(t, mowerId, mower.id())
}
