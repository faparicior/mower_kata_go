package entities

import "example.kata.local/mower/mowers/valueobject"

type Surface struct {
	surfaceXSize valueobject.SurfaceXSize
	surfaceYSize valueobject.SurfaceYSize
}

func BuildSurface(xSize valueobject.SurfaceXSize, ySize valueobject.SurfaceYSize) (Surface, error) {
	return Surface{surfaceXSize: xSize, surfaceYSize: ySize}, nil
}

func (value Surface) XSize() valueobject.SurfaceXSize {
	return value.surfaceXSize
}

func (value Surface) YSize() valueobject.SurfaceYSize {
	return value.surfaceYSize
}
