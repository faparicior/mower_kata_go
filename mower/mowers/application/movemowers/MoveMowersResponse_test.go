package movemowers

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMoveMowersResponseShouldBeBuild(t *testing.T) {
	moveMowersResponse := BuildMoveMowersResponse("response")

	if reflect.TypeOf(moveMowersResponse).String() != "movemowers.MoveMowersResponse" {
		t.Fatal(reflect.TypeOf(moveMowersResponse))
	}

	assert.Equal(t, "response", moveMowersResponse.Response())
}
