package repositories

import (
	sq "github.com/Masterminds/squirrel"
	"edb/entities/account"
	"time"
	"database/sql"
	accountError "edb/errors/account"
	"edb/errors/database"
)

type AccountRepository struct {
	Repository
}

func (repository *AccountRepository) Create(form account.CreationForm) (err error) {
	db := repository.getDB()

	_, err = sq.Insert("accounts").
		Columns("firstname", "lastname", "email", "password", "created_at", "updated_at").
		Values(form.Firstname, form.Lastname, form.Email, form.Password, time.Now().Format("2006-01-02 15:04:05"), time.Now().Format("2006-01-02 15:04:05")).
		RunWith(db).Exec()

	if err != nil {
		err = &database.Error{Message: "There is a problem inserting data"}
	}

	return
}

func (repository *AccountRepository) GetByEmail(email string) (accountObj account.Account, err error) {
	db := repository.getDB()
	accountObj = account.Account{}

	sqlErr := sq.Select("firstname, lastname, email, balance").From("accounts").Where(sq.Eq{"email": email}).
		RunWith(db).QueryRow().
		Scan(&accountObj.Firstname, &accountObj.Lastname, &accountObj.Email, &accountObj.Balance)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			err = &accountError.NotFound{Message: "Cannot find account with email: " + email}
		} else {
			err = &database.Error{Message: "There is a problem fetching data"}
		}
	}

	return
}

func (repository *AccountRepository) GetLockedForUpdateById(id uint) (accountObj account.Account, err error) {
	tx := repository.getTx()
	accountObj = account.Account{}

	sqlErr := tx.QueryRow("SELECT firstname, lastname, email, balance FROM accounts where id = ? FOR UPDATE", id).
		Scan(&accountObj.Firstname, &accountObj.Lastname, &accountObj.Email, &accountObj.Balance)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			err = &accountError.NotFound{Message: "Cannot find account with id: " + string(id)}
		} else {
			err = &database.Error{Message: "There is a problem fetching data"}
		}
	}

	return
}

func (repository *AccountRepository) UpdateBalance(id uint, balance uint) (err error) {
	tx := repository.getTx()
	_, err = sq.Update("accounts").Where(sq.Eq{"id": id}).
		Set("balance", balance).
		Set("updated_at", time.Now().Format("2006-01-02 15:04:05")).
		RunWith(tx).Exec()

	if err != nil {
		err = &database.Error{Message: "There is a problem updating data"}
	}

	return
}
