package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		PanicError(errRollback)
	} else {
		errCommit := tx.Commit()
		PanicError(errCommit)
	}
}
