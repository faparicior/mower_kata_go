package valueobject

import "errors"

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

	return MowerOrientation{}, errors.New("error")
}
