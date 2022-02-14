package valueobject

import "errors"

type MowerMovement struct {
	value string
}

func (value MowerMovement) Value() string {
	return value.value
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

	return MowerMovement{}, errors.New("error")
}
