package valueobject

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)
import "github.com/google/uuid"

func TestMowerIdShouldBeBuild(t *testing.T) {
	uuidAsString := uuid.NewString()
	mowerId, _ := BuildMowerId(uuidAsString)

	if reflect.TypeOf(mowerId).String() != "valueobject.MowerId" {
		t.Fatal(reflect.TypeOf(mowerId))
	}

	assert.Equal(t, uuidAsString, mowerId.Value())
}
