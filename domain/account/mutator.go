package account

import (
	"edb/entities/account"
	"edb/entities/transaction"
	accountError "edb/errors/account"
)

type Mutator struct {
	Logger     TransactionLogger
	Repository MutatorRepository
	DBHelper   DBHelper
}

type TransactionLogger interface {
	LogTransaction(log transaction.Log) error
}

type MutatorRepository interface {
	GetLockedForUpdateById(id uint) (account.Account, error)
	UpdateBalance(id uint, balance uint) error
}

type DBHelper interface {
	BeginTransaction() error
	CommitTransaction() error
	RollbackTransaction() error
}

func (mutator *Mutator) Mutate(accountId uint, amount int, description string) (err error) {
	err = mutator.DBHelper.BeginTransaction()
	if err != nil {
		return err
	}

	accountObj, err := mutator.Repository.GetLockedForUpdateById(accountId)
	if err != nil {
		return err
	}

	balance := int(accountObj.Balance)
	balance += amount

	if balance < 0 {
		err = &accountError.NegativeBalance{Message: "account do not have sufficient fund to perform this transaction"}
		return
	}

	transactionLog := transaction.Log{AccountId: accountId, Amount: int(amount), Description: description, Balance: uint(balance)}

	err = mutator.Logger.LogTransaction(transactionLog)
	if err != nil {
		mutator.DBHelper.RollbackTransaction()
		return err
	}

	err = mutator.Repository.UpdateBalance(accountId, uint(balance))
	if err != nil {
		mutator.DBHelper.RollbackTransaction()
		return err
	}

	return mutator.DBHelper.CommitTransaction()
}
