package account

import (
	"edb/entities/account"
	accountError "edb/errors/account"
)

type Creator struct {
	Repository CreatorRepository
}

type CreatorRepository interface {
	Create(form account.CreationForm) error
	GetByEmail(email string) (account.Account, error)
}

func (creator *Creator) Create(form account.CreationForm) (err error) {
	if len(form.Password) < 8 {
		err = &accountError.PasswordTooWeak{Message: "password length should be longer than 8 characters"}
		return
	}

	_, getErr := creator.Repository.GetByEmail(form.Email)

	if getErr == nil {
		err = &accountError.EmailTaken{Message: "Email " + form.Email + " has been taken"}
		return
	}

	err = creator.Repository.Create(form)
	return
}
