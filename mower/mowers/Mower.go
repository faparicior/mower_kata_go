package mowers

import XPosition "example.kata.local/mower/mowers/valueobject"

type Mower struct {
	XPosition.XPosition
}

func Build(xPosition XPosition.XPosition) Mower {
	return Mower{
		xPosition,
	}
}
