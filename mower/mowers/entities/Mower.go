package entities

import (
	"example.kata.local/mower/mowers/valueobject"
)

type Mower struct {
	mowerId       valueobject.MowerId
	mowerPosition valueobject.MowerPosition
}

func BuildMower(mowerId valueobject.MowerId, mowerPosition valueobject.MowerPosition) (Mower, error) {
	return Mower{
		mowerId:       mowerId,
		mowerPosition: mowerPosition,
	}, nil
}

func (value Mower) position() valueobject.MowerPosition {
	return value.mowerPosition
}

func (value Mower) id() valueobject.MowerId {
	return value.mowerId
}
