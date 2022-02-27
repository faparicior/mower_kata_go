package movemowers

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMoveOwnersCommandShouldBeBuild(t *testing.T) {
	moveMowersCommand := BuildMoveMowersCommand("filename")

	if reflect.TypeOf(moveMowersCommand).String() != "movemowers.MoveMowersCommand" {
		t.Fatal(reflect.TypeOf(moveMowersCommand))
	}

	assert.Equal(t, "filename", moveMowersCommand.filename())
}
