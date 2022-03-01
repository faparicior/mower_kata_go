package entities

import (
	"example.kata.local/mower/internal/mowers/domain/exceptions"
	valueobjects2 "example.kata.local/mower/internal/mowers/domain/valueobjects"
	"strconv"
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

func (value Mower) PositionAsString() string {
	position := value.Position()

	return strconv.Itoa(position.XPosition().Value()) +
		" " + strconv.Itoa(position.YPosition().Value()) +
		" " + position.Orientation().Value()
}

func (value Mower) id() valueobjects2.MowerId {
	return value.mowerId
}

func (value *Mower) Move(movement valueobjects2.MowerMovement, surface valueobjects2.Surface) error {
	var err error = nil

	value.mowerPosition, err = value.mowerPosition.Move(movement)

	if mowerMovementHasAnErrorByNegativeValues(err) || mowerIsOutOfBounds(value, surface) {
		return exceptions.BuildMowerOutOfBoundsError()
	}

	return nil
}

func mowerMovementHasAnErrorByNegativeValues(err error) bool {
	return nil != err
}

func mowerIsOutOfBounds(mower *Mower, surface valueobjects2.Surface) bool {
	return mower.mowerPosition.XPosition().Value() > surface.XSize().Value() ||
		mower.mowerPosition.YPosition().Value() > surface.YSize().Value()
}
