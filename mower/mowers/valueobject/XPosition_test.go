package valueobject

import (
	"encoding/json"
	"example.kata.local/mower/shared/domain/exceptions"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXPositionShouldBeBuild(t *testing.T) {
	var xPosition, err = BuildXPosition(15)

	if reflect.TypeOf(xPosition).String() != "valueobject.XPosition" {
		t.Fatal(reflect.TypeOf(xPosition))
	}

	assert.Equal(t, 15, xPosition.Value())
	//if xPosition.Value() != 15 {
	//	t.Fatal("error")
	//}

	assert.Nil(t, err)
}

func TestShouldThrownExceptionForNegativeValues(t *testing.T) {
	var xPosition, err = BuildXPosition(-1)

	assert.Equal(t, -1, xPosition.Value())
	assert.Error(t, err)
	//var dat map[string]interface{}

	var dat exceptions.Error
	var expected = exceptions.Error{
		Description:  "negative position",
		ErrorType:    "domain",
		ErrorSubtype: "InvalidMowerPosition",
	}

	err = json.Unmarshal([]byte(err.Error()), &dat)
	if err != nil {
		return
	}

	assert.EqualValues(t, dat, expected)
}

func TestShouldSumAStepForward(t *testing.T) {
	var xPosition, _ = BuildXPosition(15)

	xPosition, _ = xPosition.MoveForward(1)

	assert.Equal(t, 16, xPosition.Value())
}