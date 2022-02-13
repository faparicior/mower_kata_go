package valueobject

import "errors"

type MowerOrientation struct {
	Orientation string
}

var validOrientations = []MowerOrientation{
	{"N"},
	{"S"},
	{"E"},
	{"W"},
}

func BuildMowerOrientation(value string) (MowerOrientation, error) {

	var mowerOrientation = MowerOrientation{Orientation: value}

	for _, valid := range validOrientations {
		if valid.Orientation == value {
			return mowerOrientation, nil
		}
	}

	return mowerOrientation, errors.New("error")
}
