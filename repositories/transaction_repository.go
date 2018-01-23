package repositories

import (
	sq "github.com/Masterminds/squirrel"
	"time"
	"edb/errors/database"
	"edb/entities/transaction"
)

type TransactionRepository struct {
	Repository
}

func (repository *TransactionRepository) Create(log transaction.Log) (err error) {
	db := repository.getDB()

	_, err = sq.Insert("transactions").
		Columns("account_id", "amount", "description", "balance", "timestamp").
		Values(log.AccountId, log.Amount, log.Description, log.Balance, time.Now().Format("2006-01-02 15:04:05")).
		RunWith(db).Exec()

	if err != nil {
		err = &database.Error{Message: "There is a problem inserting data"}
	}

	return
}
