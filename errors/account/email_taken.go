package account

type EmailTaken struct {
	Message string
}

func (e *EmailTaken) Error() string {
	return e.Message
}
