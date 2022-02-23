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

func (value Mower) Position() valueobject.MowerPosition {
	return value.mowerPosition
}

func (value Mower) id() valueobject.MowerId {
	return value.mowerId
}

func (value *Mower) Move(movement valueobject.MowerMovement) {
	value.mowerPosition, _ = value.mowerPosition.Move(movement)
}
