package postgres

import (
	"context"
	"database/sql"
)

type ExecStmt func(*sql.Tx) error

type Drivers interface {
	QueryContext(ctx context.Context, sql string) (*sql.Rows, error)
	ExecContext(ctx context.Context, query string) (sql.Result, error)
	DoTransaction(ctx context.Context, fnStmt ExecStmt) error
}
