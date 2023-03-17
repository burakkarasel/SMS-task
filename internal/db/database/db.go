package db

import (
	"context"
	"database/sql"
)

// DBTX implements the database/sql functions we need
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// New Creates a new instance of Queries
func New(db DBTX) *Queries {
	return &Queries{db: db}
}

// Queries holds the DBTX interface in it, which consist of the database/sql functions we need
type Queries struct {
	db DBTX
}
