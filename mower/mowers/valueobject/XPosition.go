package XPosition

import (
	"example.kata.local/mower/mowers/exceptions"
)

type XPosition struct {
	// explicitly not public field
	value int
}

func Build(value int) (xPosition XPosition, err error) {
	if 0 > value {
		return XPosition{value: value}, exceptions.BuildInvalidMowerPosition()
	}

	return XPosition{value: value}, nil
}

func (xPosition XPosition) Value() int {
	return xPosition.value
}

func (xPosition XPosition) MoveForward(step int) (this XPosition, err error) {
	var newPosition = xPosition.value + step

	return XPosition{value: newPosition}, nil
}
