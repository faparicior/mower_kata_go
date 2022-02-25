package valueobjects

import (
	"example.kata.local/mower/mowers/domain/exceptions"
)

type YPosition struct {
	// explicitly not public field
	value int
}

func BuildYPosition(value int) (yPosition YPosition, err error) {
	const minimumYPosition = 0

	if minimumYPosition > value {
		return YPosition{}, exceptions.BuildInvalidMowerPosition()
	}

	return YPosition{value: value}, nil
}

func (yPosition YPosition) Value() int {
	return yPosition.value
}

func (yPosition YPosition) MoveForward(step int) (this YPosition, err error) {
	var newPosition = yPosition.value + step

	return YPosition{value: newPosition}, nil
}
