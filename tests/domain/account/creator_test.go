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

/*
CreatorRepository with no error return
*/

type MockErrorCreatorRepository struct{}

func (repository *MockErrorCreatorRepository) Create(form account.CreationForm) error {
	return errors.New("something went wrong")
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
