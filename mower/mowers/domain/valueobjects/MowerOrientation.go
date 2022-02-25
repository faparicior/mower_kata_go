package valueobjects

import (
	"errors"
	"example.kata.local/mower/mowers/domain/exceptions"
	"reflect"
)

type MowerOrientation struct {
	value string
}

type enumMowerOrientation struct {
	value         string
	affectsYAxis  bool
	affectsXAxis  bool
	stepDirection int
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

func (value MowerOrientation) AffectsYAxis() bool {
	return getOrientationSpecs(value).affectsYAxis
}

func (value MowerOrientation) AffectsXAxis() bool {
	return getOrientationSpecs(value).affectsXAxis
}

func (value MowerOrientation) StepMovement() int {
	return getOrientationSpecs(value).stepDirection
}

func getCurrentOrientationIndex(mowerOrientation MowerOrientation) (int, error) {
	const invalidIndex = -1

	for key, orientation := range validOrientations() {
		if orientation.value == mowerOrientation.value {
			return key, nil
		}
	}

	return invalidIndex, errors.New("orientation not valid")
}

func getOrientationSpecs(value MowerOrientation) enumMowerOrientation {
	index, _ := getCurrentOrientationIndex(value)
	orientations := validOrientations()

	return orientations[index]
}

func validOrientations() []enumMowerOrientation {
	return []enumMowerOrientation{
		{"N", true, false, 1},
		{"E", false, true, 1},
		{"S", true, false, -1},
		{"W", false, true, -1},
	}
}
