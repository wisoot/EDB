package account

import (
	accountDomain "edb/domain/account"
	"edb/entities/account"
	"edb/entities/transaction"
	accountError "edb/errors/account"
	"testing"
)

type MockDbHelper struct{}

func (repository *MockDbHelper) BeginTransaction() error {
	return nil
}

func (repository *MockDbHelper) CommitTransaction() error {
	return nil
}

func (repository *MockDbHelper) RollbackTransaction() error {
	return nil
}

/*
CreditorRepository with no error return
*/

type MockMutatorRepository struct{}

func (repository *MockMutatorRepository) GetLockedForUpdateById(id uint) (accountObj account.Account, err error) {
	accountObj = account.Account{Id: 1, Firstname: "Obi-Wan", Lastname: "Kenobe", Email: "example@example.com", Balance: 140}
	return
}

func (repository *MockMutatorRepository) UpdateBalance(id uint, balance uint) error {
	return nil
}

/*
CreditorRepository with no account return
*/

type MockMutatorRepositoryNoAccount struct{}

func (repository *MockMutatorRepositoryNoAccount) GetLockedForUpdateById(id uint) (accountObj account.Account, err error) {
	err = &accountError.NotFound{Message: "account not found"}
	return
}

func (repository *MockMutatorRepositoryNoAccount) UpdateBalance(id uint, balance uint) error {
	return nil
}

/*
TransactionLogger with no error return
*/

type MockTransactionLogger struct{}

func (logger *MockTransactionLogger) LogTransaction(log transaction.Log) error {
	return nil
}

func TestMutate(t *testing.T) {
	mutator := accountDomain.Mutator{Logger: &MockTransactionLogger{}, DBHelper: &MockDbHelper{}, Repository: &MockMutatorRepository{}}

	err := mutator.Mutate(1, 1000, "Deposit $10")

	if err != nil {
		t.Error("Expect err to be nil got something else")
	}
}

func TestMutateCannotFindAccount(t *testing.T) {
	mutator := accountDomain.Mutator{Logger: &MockTransactionLogger{}, DBHelper: &MockDbHelper{}, Repository: &MockMutatorRepositoryNoAccount{}}

	err := mutator.Mutate(1, 1000, "Deposit $10")

	if err == nil {
		t.Error("Expect error got nil")
	} else {
		_, ok := err.(*accountError.NotFound)

		if !ok {
			t.Error("Expected NotFound error got something else")
		}
	}
}

func TestMutateResultInNegativeBalance(t *testing.T) {
	mutator := accountDomain.Mutator{Logger: &MockTransactionLogger{}, DBHelper: &MockDbHelper{}, Repository: &MockMutatorRepository{}}

	err := mutator.Mutate(1, -1000, "Withdraw $10")

	if err == nil {
		t.Error("Expect error got nil")
	} else {
		_, ok := err.(*accountError.NegativeBalance)

		if !ok {
			t.Error("Expected NegativeBalance error got something else")
		}
	}
}
