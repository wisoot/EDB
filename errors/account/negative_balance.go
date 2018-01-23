package account

type NegativeBalance struct {
	Message string
}

func (e *NegativeBalance) Error() string {
	return e.Message
}
