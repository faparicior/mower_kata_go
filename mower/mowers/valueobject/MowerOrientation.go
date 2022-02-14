package valueobject

import "errors"

type MowerOrientation struct {
	orientation string
}

var validOrientations = []MowerOrientation{
	{"N"},
	{"S"},
	{"E"},
	{"W"},
}

func BuildMowerOrientation(value string) (MowerOrientation, error) {

	var mowerOrientation = MowerOrientation{orientation: value}

	for _, validValues := range validOrientations {
		if validValues.orientation == value {
			return mowerOrientation, nil
		}
	}

	return mowerOrientation, errors.New("error")
}
