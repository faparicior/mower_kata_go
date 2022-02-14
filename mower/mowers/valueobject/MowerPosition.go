package valueobject

type MowerPosition struct {
	xMowerPosition XPosition
	yMowerPosition YPosition
}

func BuildMowerPosition(xPosition XPosition, yPosition YPosition) (MowerPosition, error) {
	return MowerPosition{
		xMowerPosition: xPosition,
		yMowerPosition: yPosition,
	}, nil
}

func (value *MowerPosition) XPosition() XPosition {
	return value.xMowerPosition
}

func (value *MowerPosition) YPosition() YPosition {
	return value.yMowerPosition
}
