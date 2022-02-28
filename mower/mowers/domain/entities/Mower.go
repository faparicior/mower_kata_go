package entities

import (
	"example.kata.local/mower/mowers/domain/exceptions"
	"example.kata.local/mower/mowers/domain/valueobjects"
	"strconv"
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

func (value Mower) PositionAsString() string {
	position := value.Position()

	return strconv.Itoa(position.XPosition().Value()) +
		" " + strconv.Itoa(position.YPosition().Value()) +
		" " + position.Orientation().Value()
}

func (value Mower) id() valueobjects.MowerId {
	return value.mowerId
}

func (value *Mower) Move(movement valueobjects.MowerMovement, surface valueobjects.Surface) error {
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

func mowerIsOutOfBounds(mower *Mower, surface valueobjects.Surface) bool {
	return mower.mowerPosition.XPosition().Value() > surface.XSize().Value() ||
		mower.mowerPosition.YPosition().Value() > surface.YSize().Value()
}
