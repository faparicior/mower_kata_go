package exceptions

const ErrorType = "domain"

type Error struct {
	Description  string
	ErrorType    string
	ErrorSubtype string
}
