package valueobject

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMowerPositionShouldBeBuild(t *testing.T) {
	xPosition, _ := BuildXPosition(15)
	yPosition, _ := BuildYPosition(8)

	mowerPosition, _ := BuildMowerPosition(xPosition, yPosition)

	if reflect.TypeOf(mowerPosition).String() != "valueobject.MowerPosition" {
		t.Fatal(reflect.TypeOf(xPosition))
	}

	assert.Equal(t, xPosition, mowerPosition.XPosition())
	assert.Equal(t, yPosition, mowerPosition.YPosition())
}
