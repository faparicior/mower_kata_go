package entities

import (
	valueobjects2 "example.kata.local/mower/mowers/domain/valueobjects"
)

type Mower struct {
	mowerId       valueobjects2.MowerId
	mowerPosition valueobjects2.MowerPosition
}

func BuildMower(mowerId valueobjects2.MowerId, mowerPosition valueobjects2.MowerPosition) (Mower, error) {
	return Mower{
		mowerId:       mowerId,
		mowerPosition: mowerPosition,
	}, nil
}

func (value Mower) Position() valueobjects2.MowerPosition {
	return value.mowerPosition
}

func (value Mower) id() valueobjects2.MowerId {
	return value.mowerId
}

func (value *Mower) Move(movement valueobjects2.MowerMovement) {
	value.mowerPosition, _ = value.mowerPosition.Move(movement)
}
