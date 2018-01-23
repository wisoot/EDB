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
}

func (creator *Creator) Create(form account.CreationForm) (err error) {
	if len(form.Password) < 8 {
		err = &accountError.PasswordTooWeak{Message: "password length should be longer than 8 characters"}
		return
	}

	err = creator.Repository.Create(form)
	return
}
