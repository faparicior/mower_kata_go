package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoveMowersUseCaseShouldProcessInputFileAndOutputResults(t *testing.T) {
	var output bytes.Buffer
	args := []string{"mowers/infrastructure/testdata/instructions.txt"}

	submain(&output, args)

	assert.Equal(t, "1 3 N\n5 1 E\n", output.String())
}
