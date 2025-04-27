package Database

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	Connect()
	Close() bool
	Reconnect()
	Version() (version interface{}, error error)
	Get(dest interface{}, query string, args ...interface{}) error
	Execute(query string, args ...interface{}) (sql.Result, error)
	MapExecute(query string, arg interface{}) (sql.Result, error)
	TransactionBegin() *sqlx.Tx
}

func NewDB(db DB) DB {
	return db
}
