package valueobjects

import (
	"example.kata.local/mower/internal/mowers/domain/exceptions"
)

type SurfaceXSize struct {
	// explicitly not public field
	value int
}

func BuildSurfaceXSize(value int) (surfaceXSize SurfaceXSize, err error) {
	const minimumSurfaceSize = 0

	if minimumSurfaceSize > value {
		return SurfaceXSize{}, exceptions.BuildInvalidSurfaceSize()
	}

	return SurfaceXSize{value: value}, nil
}

func (surfaceXSize SurfaceXSize) Value() int {
	return surfaceXSize.value
}
