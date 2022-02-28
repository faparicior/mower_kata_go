package infrastructure

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

var genericInstructions Instructions

func setUpTest() {
	provider := BuildFlatFileInstructionsProvider()
	genericInstructions, _ = provider.Load("testdata/instructions.txt")
}

func TestFlatFileInstructionsProviderShouldBeBuild(t *testing.T) {
	provider := BuildFlatFileInstructionsProvider()

	if reflect.TypeOf(provider).String() != "infrastructure.Instructions" {
		t.Fatal(reflect.TypeOf(provider))
	}

	expectedProvider := Instructions{
		surface:      "",
		instructions: nil,
	}

	assert.Equal(t, expectedProvider, provider)
}

func TestFlatFileInstructionsProviderShouldLoadFile(t *testing.T) {
	provider := BuildFlatFileInstructionsProvider()
	instructions, _ := provider.Load("testdata/instructions.txt")

	surface, _ := instructions.Surface()

	assert.Equal(t, 5, surface.XSize().Value())
}

func TestFlatFileInstructionsProviderShouldLoadSurface(t *testing.T) {
	setUpTest()
	surface, _ := genericInstructions.Surface()

	assert.Equal(t, 5, surface.XSize().Value())
	assert.Equal(t, 5, surface.YSize().Value())
}

func TestFlatFileInstructionsProviderShouldReturnMowerPositionByIndex(t *testing.T) {
	setUpTest()

	mowerPosition, _ := genericInstructions.MowerPosition(0)

	assert.Equal(t, 1, mowerPosition.XPosition().Value())
	assert.Equal(t, 2, mowerPosition.YPosition().Value())
	assert.Equal(t, "N", mowerPosition.Orientation().Value())
}

func TestFlatFileInstructionsProviderShouldReturnMowerMovementsByIndex(t *testing.T) {
	setUpTest()

	mowerMovements := genericInstructions.MowerMovements(1)

	expectedResult := strings.Split("FFRFFRFRRF", "")

	assert.Equal(t, expectedResult, mowerMovements)
}
