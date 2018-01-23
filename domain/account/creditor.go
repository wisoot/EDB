package account

type Creditor struct {
	Mutator
}

func (creditor *Creditor) Credit(accountId uint, amount uint, description string) error {
	return creditor.Mutate(accountId, int(amount), description)
}
