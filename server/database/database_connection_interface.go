package database

import (
	"github.com/jmoiron/sqlx"
	"database/sql"
)

type DatabaseConnection interface {
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	MustBegin() *sqlx.Tx
	Preparex(query string) (*sqlx.Stmt, error)
	sqlx.Ext
	sqlx.Preparer
	QueryRow(query string, args ...interface{}) *sql.Row
	Beginx() (*sqlx.Tx, error)
}