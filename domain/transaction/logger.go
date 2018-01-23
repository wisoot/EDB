package transaction

import "edb/entities/transaction"

type Logger struct {
	Repository LoggerRepository
}

type LoggerRepository interface {
	Create(log transaction.Log) error
}

func (logger *Logger) LogTransaction(log transaction.Log) error {
	return logger.Repository.Create(log)
}
