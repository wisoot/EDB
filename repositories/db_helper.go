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
	err := tx.Commit()
	dbHelper.clearTx()
	return err
}

func (dbHelper *DBHelper) RollbackTransaction() error {
	tx := dbHelper.getTx()
	err := tx.Rollback()
	dbHelper.clearTx()
	return err
}
