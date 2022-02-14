package valueobject

import (
	"example.kata.local/mower/mowers/exceptions"
)

type MowerOrientation struct {
	value string
}

func BuildMowerOrientation(value string) (MowerOrientation, error) {
	var validOrientations = []MowerOrientation{
		{"N"},
		{"S"},
		{"E"},
		{"W"},
	}

	for _, validValues := range validOrientations {
		if validValues.value == value {
			return MowerOrientation{value: value}, nil
		}
	}

	return MowerOrientation{}, exceptions.BuildInvalidMowerOrientation()
}
