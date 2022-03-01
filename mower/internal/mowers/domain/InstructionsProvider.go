package domain

import (
	valueobjects2 "example.kata.local/mower/internal/mowers/domain/valueobjects"
)

type InstructionsProvider interface {
	Surface() (valueobjects2.Surface, error)
	MowerPosition(index int) (valueobjects2.MowerPosition, error)
	MowerMovements(index int) []valueobjects2.MowerMovement
	CountMowers() int
}
