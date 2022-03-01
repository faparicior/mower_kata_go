package domain

import "example.kata.local/mower/mowers/domain/valueobjects"

type InstructionsProvider interface {
	Surface() (valueobjects.Surface, error)
	MowerPosition(index int) (valueobjects.MowerPosition, error)
	MowerMovements(index int) []valueobjects.MowerMovement
	CountMowers() int
}
