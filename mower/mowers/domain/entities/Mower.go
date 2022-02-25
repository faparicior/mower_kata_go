package entities

import (
	"example.kata.local/mower/mowers/domain/valueobjects"
)

type Mower struct {
	mowerId       valueobjects.MowerId
	mowerPosition valueobjects.MowerPosition
}

func BuildMower(mowerId valueobjects.MowerId, mowerPosition valueobjects.MowerPosition) (Mower, error) {
	return Mower{
		mowerId:       mowerId,
		mowerPosition: mowerPosition,
	}, nil
}

func (value Mower) Position() valueobjects.MowerPosition {
	return value.mowerPosition
}

func (value Mower) id() valueobjects.MowerId {
	return value.mowerId
}

func (value *Mower) Move(movement valueobjects.MowerMovement) {
	value.mowerPosition, _ = value.mowerPosition.Move(movement)
}
