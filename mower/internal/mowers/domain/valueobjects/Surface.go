package valueobjects

type Surface struct {
	surfaceXSize SurfaceXSize
	surfaceYSize SurfaceYSize
}

func BuildSurface(xSize SurfaceXSize, ySize SurfaceYSize) (Surface, error) {
	return Surface{surfaceXSize: xSize, surfaceYSize: ySize}, nil
}

func (value Surface) XSize() SurfaceXSize {
	return value.surfaceXSize
}

func (value Surface) YSize() SurfaceYSize {
	return value.surfaceYSize
}
