package types

import (
	"database/sql"
	"github.com/davecgh/go-spew/spew"
	"github.com/jmoiron/sqlx"
)

type LoggedDB struct {
	sqlx.Ext
}

func (d LoggedDB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	println(query)
	spew.Dump(args)
	return d.Ext.Query(query, args...)
}

func (d LoggedDB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	println(query)
	spew.Dump(args)
	return d.Ext.Queryx(query, args...)
}

func (d LoggedDB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	println(query)
	spew.Dump(args)
	return d.Ext.QueryRowx(query, args...)
}

func (d LoggedDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	println(query)
	spew.Dump(args)
	return d.Ext.Exec(query, args...)
}
