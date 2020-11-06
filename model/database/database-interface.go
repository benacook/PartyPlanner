package database

import "database/sql"

type Database interface {
	Close()
	Exec(query string, args... interface{}) (sql.Result, error)
	Query(query string, args... interface{}) (*sql.Rows, error)
}
