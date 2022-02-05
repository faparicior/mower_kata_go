package mower

import (
	"reflect"
	"testing"

	"example.kata.local/mower/mowers/valueobject"
)

func TestBuild(t *testing.T) {
	var xPosition = XPosition.Build(15)

	if reflect.TypeOf(xPosition).String() != "XPosition.Instance" {
		t.Fatal(reflect.TypeOf(xPosition))
	}

	if xPosition.Value() != 15 {
		t.Fatal("error")
	}
}
