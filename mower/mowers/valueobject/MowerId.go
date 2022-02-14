package valueobject

type MowerId struct {
	id string
}

func BuildMowerId(value string) (MowerId, error) {
	return MowerId{id: value}, nil
}

func (value MowerId) Value() string {
	return value.id
}
