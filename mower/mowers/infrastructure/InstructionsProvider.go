package infrastructure

import "example.kata.local/mower/mowers/domain/valueobjects"

type InstructionsProvider interface {
	Load() Instructions
	Surface() valueobjects.Surface
}

type Instructions struct {
	surface      string
	instructions []string
}
