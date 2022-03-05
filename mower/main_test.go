package main

import (
	"bytes"
	"example.kata.local/mower/internal/mowers/ui"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveMowersUseCaseShouldProcessInputFileAndOutputResults(t *testing.T) {
	var output bytes.Buffer
	args := []string{"internal/mowers/infrastructure/testdata/instructions.txt"}

	ui.Cmd(&output, args)

	assert.Equal(t, "1 3 N\n5 1 E\n", output.String())
}
