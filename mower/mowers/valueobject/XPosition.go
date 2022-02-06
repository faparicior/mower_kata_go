package XPosition

import "errors"

type Instance struct {
	// explicitly not public field
	value int
}

func Build(value int) (this Instance, err error) {
	if 0 > value {
		return Instance{value: value}, errors.New("negative position")
	}

	return Instance{value: value}, nil
}

func (instance Instance) Value() int {
	return instance.value
}

func (instance Instance) MoveForward(step int) (this Instance, err error) {
	var newPosition = instance.value + step

	return Instance{value: newPosition}, nil
}
