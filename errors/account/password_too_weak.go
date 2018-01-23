package account

type PasswordTooWeak struct {
	Message string
}

func (e *PasswordTooWeak) Error() string {
	return e.Message
}
