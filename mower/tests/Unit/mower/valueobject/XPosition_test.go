package mower

import (
	"reflect"
	"testing"

	"example.kata.local/mower/mowers/valueobject"
	"github.com/stretchr/testify/assert"
)

func TestShouldBeBuild(t *testing.T) {
	var xPosition, err = XPosition.Build(15)

	if reflect.TypeOf(xPosition).String() != "XPosition.Instance" {
		t.Fatal(reflect.TypeOf(xPosition))
	}

	assert.Equal(t, 15, xPosition.Value())
	//if xPosition.Value() != 15 {
	//	t.Fatal("error")
	//}

	assert.Nil(t, err)
}

func TestShouldThrownExceptionForNegativeValues(t *testing.T) {
	var xPosition, err = XPosition.Build(-1)

	assert.Equal(t, -1, xPosition.Value())
	assert.Error(t, err)
	assert.EqualError(t, err, "negative position")
}

func TestShouldSumAStepForward(t *testing.T) {
	var xPosition, _ = XPosition.Build(15)

	xPosition, _ = xPosition.MoveForward(1)

	assert.Equal(t, 16, xPosition.Value())
}
