package account

type Debitor struct {
	Mutator
}

func (debitor *Debitor) Debit(accountId uint, amount uint, description string) error {
	return debitor.Mutate(accountId, -1*int(amount), description)
}
