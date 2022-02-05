package XPosition

type Instance struct {
	// explicitly not public field
	value int
}

func Build(value int) Instance {
	return Instance{value: value}
}

func (value Instance) Value() int {
	return value.value
}
