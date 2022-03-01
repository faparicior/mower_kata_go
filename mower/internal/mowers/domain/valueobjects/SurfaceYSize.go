package valueobjects

import (
	"example.kata.local/mower/internal/mowers/domain/exceptions"
)

type SurfaceYSize struct {
	// explicitly not public field
	value int
}

func BuildSurfaceYSize(value int) (surfaceYSize SurfaceYSize, err error) {
	const minimumSurfaceSize = 0

	if minimumSurfaceSize > value {
		return SurfaceYSize{}, exceptions.BuildInvalidSurfaceSize()
	}

	return SurfaceYSize{value: value}, nil
}

func (surfaceYSize SurfaceYSize) Value() int {
	return surfaceYSize.value
}
