package account

import "edb/entities/account"

type Creator struct {
	Repository CreatorRepository
}

type CreatorRepository interface {
	Create(form account.CreationForm) error
}

func (creator *Creator) Create(form account.CreationForm) (err error) {
	err = creator.Repository.Create(form)
	return
}
