package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		PanicIfError("Rollback get error", errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		PanicIfError("Commit get error", errorCommit)
	}
}
