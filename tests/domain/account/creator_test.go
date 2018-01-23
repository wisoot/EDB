package account

import (
	accountDomain "edb/domain/account"
	"edb/entities/account"
	accountError "edb/errors/account"
	"errors"
	"testing"
)

/*
CreatorRepository with no error return
*/

type MockCreatorRepository struct{}

func (repository *MockCreatorRepository) Create(form account.CreationForm) error {
	return nil
}

func (repository *MockCreatorRepository) GetByEmail(email string) (accountObj account.Account, err error) {
	err = &accountError.NotFound{Message: "account not found"}
	return
}

/*
CreatorRepository with no error return
*/

type MockErrorCreatorRepository struct{}

func (repository *MockErrorCreatorRepository) Create(form account.CreationForm) error {
	return errors.New("something went wrong")
}

func (repository *MockErrorCreatorRepository) GetByEmail(email string) (accountObj account.Account, err error) {
	err = &accountError.NotFound{Message: "account not found"}
	return
}

/*
CreatorRepository with existing account taken the email
*/

type MockCreatorRepositoryWithExistingAccount struct{}

func (repository *MockCreatorRepositoryWithExistingAccount) Create(form account.CreationForm) error {
	return nil
}

func (repository *MockCreatorRepositoryWithExistingAccount) GetByEmail(email string) (accountObj account.Account, err error) {
	accountObj = account.Account{Id: 1, Firstname: "Obi-Wan", Lastname: "Kenobe", Email: "example@example.com", Balance: 0}
	return
}

/*
Test cases
*/

func TestCreate(t *testing.T) {
	form := account.CreationForm{Firstname: "Obi-Wan", Lastname: "Kenobe", Email: "example@example.com", Password: "abcd1234"}

	creator := accountDomain.Creator{Repository: &MockCreatorRepository{}}
	err := creator.Create(form)

	if err != nil {
		t.Error("Expected nil got ", err)
	}
}

func TestCreateFail(t *testing.T) {
	form := account.CreationForm{Firstname: "Obi-Wan", Lastname: "Kenobe", Email: "example@example.com", Password: "abcd1234"}

	creator := accountDomain.Creator{Repository: &MockErrorCreatorRepository{}}
	err := creator.Create(form)

	if err == nil {
		t.Error("Expected error got nil")
	}
}

func TestCreatePasswordTooWeak(t *testing.T) {
	form := account.CreationForm{Firstname: "Obi-Wan", Lastname: "Kenobe", Email: "example@example.com", Password: "abc123"}

	creator := accountDomain.Creator{Repository: &MockErrorCreatorRepository{}}
	err := creator.Create(form)

	if err == nil {
		t.Error("Expected error got nil")
	} else {
		_, ok := err.(*accountError.PasswordTooWeak)

		if !ok {
			t.Error("Expected PasswordIsTooWeak error got something else")
		}
	}
}

func TestCreateEmailTaken(t *testing.T) {
	form := account.CreationForm{Firstname: "Obi-Wan", Lastname: "Kenobe", Email: "example@example.com", Password: "abcd1234"}

	creator := accountDomain.Creator{Repository: &MockCreatorRepositoryWithExistingAccount{}}
	err := creator.Create(form)

	if err == nil {
		t.Error("Expected error got nil")
	} else {
		_, ok := err.(*accountError.EmailTaken)

		if !ok {
			t.Error("Expected EmailTaken error got something else")
		}
	}
}
