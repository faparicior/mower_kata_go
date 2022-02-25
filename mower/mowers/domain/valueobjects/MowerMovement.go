package valueobjects

import (
	"example.kata.local/mower/mowers/domain/exceptions"
)

type MowerMovement struct {
	value string
}

func (value MowerMovement) Value() string {
	return value.value
}

func (value MowerMovement) isClockWise() bool {
	return "R" == value.value
}

func (value MowerMovement) isCounterClockWise() bool {
	return "L" == value.value
}

func (value MowerMovement) isForward() bool {
	return "F" == value.value
}

func BuildMowerMovement(value string) (MowerMovement, error) {
	var validMovements = []MowerMovement{
		{"L"},
		{"R"},
		{"F"},
	}

	for _, validValues := range validMovements {
		if validValues.value == value {
			return MowerMovement{value: value}, nil
		}
	}

	return MowerMovement{}, exceptions.BuildInvalidMowerMovement()
}
