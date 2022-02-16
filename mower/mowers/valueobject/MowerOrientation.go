package valueobject

import (
	"errors"
	"example.kata.local/mower/mowers/exceptions"
	"reflect"
)

type MowerOrientation struct {
	value string
}

func BuildMowerOrientation(value string) (MowerOrientation, error) {
	validOrientations := validOrientations()

	for _, validValues := range validOrientations {
		if validValues.value == value {
			return MowerOrientation{value: value}, nil
		}
	}

	return MowerOrientation{}, exceptions.BuildInvalidMowerOrientation()
}

func (value MowerOrientation) ChangeOrientation(movement MowerMovement) MowerOrientation {
	const arrayLengthCompensation int = 1
	const firstElement int = 0
	const compassStep int = 1

	validOrientations := validOrientations()
	s := reflect.ValueOf(validOrientations)
	compassElementsLength := s.Len()

	currentOrientation := value

	if movement.isClockWise() {
		currentOrientationIndex, _ := getCurrentOrientationIndex(currentOrientation)

		if compassElementsLength == currentOrientationIndex+arrayLengthCompensation {
			newOrientation, _ := BuildMowerOrientation(validOrientations[firstElement].value)
			return newOrientation
		}

		newOrientation, _ := BuildMowerOrientation(validOrientations[currentOrientationIndex+compassStep].value)

		return newOrientation
	}

	if movement.isCounterClockWise() {
		currentOrientationIndex, _ := getCurrentOrientationIndex(currentOrientation)

		if currentOrientationIndex == firstElement {
			newOrientation, _ := BuildMowerOrientation(validOrientations[compassElementsLength-arrayLengthCompensation].value)
			return newOrientation
		}

		newOrientation, _ := BuildMowerOrientation(validOrientations[currentOrientationIndex-compassStep].value)

		return newOrientation
	}

	return MowerOrientation{
		value: value.value,
	}
}

func getCurrentOrientationIndex(mowerOrientation MowerOrientation) (int, error) {
	const invalidIndex = -1

	for key, orientation := range validOrientations() {
		if orientation == mowerOrientation {
			return key, nil
		}
	}

	return invalidIndex, errors.New("orientation not valid")
}

func validOrientations() []MowerOrientation {
	return []MowerOrientation{
		{"N"},
		{"E"},
		{"S"},
		{"W"},
	}
}
