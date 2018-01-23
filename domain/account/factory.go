package account

import (
	"edb/repositories"
	"edb/domain/transaction"
)

func MakeCreator() Creator {
	return Creator{&repositories.AccountRepository{}}
}

func MakeCreditor() Creditor {
	return Creditor{MakeMutator()}
}

func MakeDebitor() Debitor {
	return Debitor{MakeMutator()}
}

func MakeMutator() Mutator {
	logger := transaction.MakeLogger()
	return Mutator{&logger, &repositories.AccountRepository{}, &repositories.DBHelper{}}
}