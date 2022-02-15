package entities

import (
	"example.kata.local/mower/mowers/valueobject"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSurfaceShouldBeBuild(t *testing.T) {
	xSize, _ := valueobject.BuildSurfaceXSize(10)
	ySize, _ := valueobject.BuildSurfaceYSize(10)

	surface, err := BuildSurface(xSize, ySize)

	if reflect.TypeOf(surface).String() != "entities.Surface" {
		t.Fatal(reflect.TypeOf(surface))
	}

	assert.Equal(t, xSize, surface.XSize())
	assert.Equal(t, ySize, surface.YSize())

	assert.Nil(t, err)
}
