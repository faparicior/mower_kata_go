package valueobjects

import (
	"example.kata.local/mower/mowers/domain/exceptions"
)

type XPosition struct {
	// explicitly not public field
	value int
}

func BuildXPosition(value int) (xPosition XPosition, err error) {
	const minimumXPosition = 0

	if minimumXPosition > value {
		return XPosition{}, exceptions.BuildInvalidMowerPosition()
	}

	return XPosition{value: value}, nil
}

func (xPosition XPosition) Value() int {
	return xPosition.value
}

func (xPosition XPosition) MoveForward(step int) (this XPosition, err error) {
	var newPosition = xPosition.value + step

	return BuildXPosition(newPosition)
}
