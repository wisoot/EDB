package transaction

import "edb/repositories"

func MakeLogger() Logger {
	return Logger{&repositories.TransactionRepository{}}
}
