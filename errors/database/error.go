package database

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
