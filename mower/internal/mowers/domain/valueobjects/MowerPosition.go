package valueobjects

type MowerPosition struct {
	xMowerPosition XPosition
	yMowerPosition YPosition
	orientation    MowerOrientation
}

func BuildMowerPosition(xPosition XPosition, yPosition YPosition, orientation MowerOrientation) (MowerPosition, error) {
	return MowerPosition{
		xMowerPosition: xPosition,
		yMowerPosition: yPosition,
		orientation:    orientation,
	}, nil
}

func (value *MowerPosition) XPosition() XPosition {
	return value.xMowerPosition
}

func (value *MowerPosition) YPosition() YPosition {
	return value.yMowerPosition
}

func (value *MowerPosition) Orientation() MowerOrientation {
	return value.orientation
}

func (value *MowerPosition) Move(movement MowerMovement) (MowerPosition, error) {

	if movement.isForward() {
		if value.Orientation().AffectsYAxis() {
			yPosition, err := value.YPosition().MoveForward(value.Orientation().StepMovement())

			if nil != err {
				return MowerPosition{}, err
			}
			return BuildMowerPosition(value.XPosition(), yPosition, value.Orientation())
		}
		if value.Orientation().AffectsXAxis() {
			xPosition, err := value.XPosition().MoveForward(value.Orientation().StepMovement())

			if nil != err {
				return MowerPosition{}, err
			}
			return BuildMowerPosition(xPosition, value.YPosition(), value.Orientation())
		}
	}

	newOrientation := value.Orientation().ChangeOrientation(movement)
	return BuildMowerPosition(value.XPosition(), value.YPosition(), newOrientation)
}
