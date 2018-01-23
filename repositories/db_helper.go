package repositories

type DBHelper struct {
	Repository
}

func (dbHelper *DBHelper) BeginTransaction() error {
	dbHelper.getTx()

	return nil
}

func (dbHelper *DBHelper) CommitTransaction() error {
	tx := dbHelper.getTx()
	return tx.Commit()
}

func (dbHelper *DBHelper) RollbackTransaction() error {
	tx := dbHelper.getTx()
	return tx.Rollback()
}
