package infrastructure

import (
	"example.kata.local/mower/mowers/domain/valueobjects"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

var genericInstructions FlatFileInstructionsProvider

func setUpTest() {
	provider := BuildFlatFileInstructionsProvider()
	genericInstructions, _ = provider.Load("testdata/instructions.txt")
}

func TestFlatFileInstructionsProviderShouldBeBuild(t *testing.T) {
	provider := BuildFlatFileInstructionsProvider()

	if reflect.TypeOf(provider).String() != "infrastructure.FlatFileInstructionsProvider" {
		t.Fatal(reflect.TypeOf(provider))
	}

	expectedProvider := FlatFileInstructionsProvider{}

	assert.Equal(t, expectedProvider, provider)
}

func TestFlatFileInstructionsProviderShouldLoadFile(t *testing.T) {
	provider := BuildFlatFileInstructionsProvider()
	instructions, _ := provider.Load("testdata/instructions.txt")

	assert.NotNil(t, instructions.filename)
	assert.NotNil(t, instructions.surface)
	assert.NotNil(t, instructions.movements)
	assert.NotNil(t, instructions.positions)
}

func TestFlatFileInstructionsProviderShouldReturnSurface(t *testing.T) {
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

	expectedMovements := strings.Split("FFRFFRFRRF", "")
	var expectedResult []valueobjects.MowerMovement

	for _, value := range expectedMovements {
		movement, _ := valueobjects.BuildMowerMovement(value)
		expectedResult = append(expectedResult, movement)
	}

	assert.Equal(t, expectedResult, mowerMovements)
}
