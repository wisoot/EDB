package account

type PasswordHashedFailed struct {
	Message string
}

func (e *PasswordHashedFailed) Error() string {
	return e.Message
}
