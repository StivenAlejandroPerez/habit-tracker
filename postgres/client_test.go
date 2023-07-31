package postgres

import (
	"context"
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPostgres_QueryContext(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		query string
	}
	type test struct {
		name    string
		fields  fields
		args    args
		want    func(t assert.TestingT, i2 ...interface{}) bool
		wantErr assert.ErrorAssertionFunc
	}

	tests := []test{
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM test WHERE id = 1;`)).WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(1),
			)

			return test{
				name: "Success",
				fields: fields{
					db: db,
				},
				args: args{
					ctx:   context.Background(),
					query: `SELECT * FROM test WHERE id = 1;`,
				},
				want: func(t assert.TestingT, i2 ...interface{}) bool {
					return assert.NotNil(t, i2)
				},
				wantErr: assert.NoError,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM test WHERE id = 1;`)).WillReturnError(assert.AnError)

			return test{
				name: "Error",
				fields: fields{
					db: db,
				},
				args: args{
					ctx:   context.Background(),
					query: `SELECT * FROM test WHERE id = 1;`,
				},
				want: func(t assert.TestingT, i2 ...interface{}) bool {
					return assert.NotNil(t, i2)
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				p := &Postgres{
					db: tt.fields.db,
				}
				got, err := p.QueryContext(tt.args.ctx, tt.args.query)

				tt.wantErr(t, err)
				tt.want(t, got)
			},
		)
	}
}

func TestPostgres_ExecContext(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx   context.Context
		query string
	}
	type test struct {
		name    string
		fields  fields
		args    args
		want    func(t assert.TestingT, i2 ...interface{}) bool
		wantErr assert.ErrorAssertionFunc
	}

	tests := []test{
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO test (id, name) VALUES (1, "test");`)).WillReturnResult(
				sqlmock.NewResult(1, 1),
			)

			return test{
				name: "Success",
				fields: fields{
					db: db,
				},
				args: args{
					ctx:   context.Background(),
					query: `INSERT INTO test (id, name) VALUES (1, "test");`,
				},
				want: func(t assert.TestingT, i2 ...interface{}) bool {
					return assert.NotNil(t, i2)
				},
				wantErr: assert.NoError,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO test (id, name) VALUES (1, "test");`)).
				WillReturnError(assert.AnError)

			return test{
				name: "Error",
				fields: fields{
					db: db,
				},
				args: args{
					ctx:   context.Background(),
					query: `INSERT INTO test (id, name) VALUES (1, "test");`,
				},
				want: func(t assert.TestingT, i2 ...interface{}) bool {
					return assert.NotNil(t, i2)
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				p := &Postgres{
					db: tt.fields.db,
				}
				got, err := p.ExecContext(tt.args.ctx, tt.args.query)

				tt.wantErr(t, err)
				tt.want(t, got)
			},
		)
	}
}

func TestPostgres_DoTransaction(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		ctx    context.Context
		fnStmt ExecStmt
	}
	type test struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}

	tests := []test{
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO test (id, name) VALUES (1, "test");`)).WillReturnResult(
				sqlmock.NewResult(1, 1),
			)
			mock.ExpectCommit()

			return test{
				name: "Success",
				fields: fields{
					db: db,
				},
				args: args{
					ctx: context.Background(),
					fnStmt: func(tx *sql.Tx) error {
						_, err = tx.ExecContext(context.Background(), `INSERT INTO test (id, name) VALUES (1, "test");`)

						return err
					},
				},
				wantErr: assert.NoError,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectBegin().WillReturnError(assert.AnError)

			return test{
				name: "ErrorBeginTx",
				fields: fields{
					db: db,
				},
				args: args{
					ctx: context.Background(),
				},
				wantErr: assert.Error,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO test (id, name) VALUES (1, "test");`)).WillReturnError(assert.AnError)
			mock.ExpectRollback()

			return test{
				name: "Error_statement",
				fields: fields{
					db: db,
				},
				args: args{
					ctx: context.Background(),
					fnStmt: func(tx *sql.Tx) error {
						_, err = tx.ExecContext(context.Background(), `INSERT INTO test (id, name) VALUES (1, "test");`)

						return err
					},
				},
				wantErr: assert.Error,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO test (id, name) VALUES (1, "test");`)).WillReturnError(assert.AnError)
			mock.ExpectRollback().WillReturnError(assert.AnError)

			return test{
				name: "ErrorRollback",
				fields: fields{
					db: db,
				},
				args: args{
					ctx: context.Background(),
					fnStmt: func(tx *sql.Tx) error {
						_, err = tx.ExecContext(context.Background(), `INSERT INTO test (id, name) VALUES (1, "test");`)

						return err
					},
				},
				wantErr: assert.Error,
			}
		}(),
		func() test {
			db, mock, err := sqlmock.New()
			assert.NoError(t, err)

			mock.ExpectBegin()
			mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO test (id, name) VALUES (1, "test");`)).WillReturnResult(
				sqlmock.NewResult(1, 1),
			)
			mock.ExpectCommit().WillReturnError(assert.AnError)

			return test{
				name: "ErrorCommit",
				fields: fields{
					db: db,
				},
				args: args{
					ctx: context.Background(),
					fnStmt: func(tx *sql.Tx) error {
						_, err = tx.ExecContext(context.Background(), `INSERT INTO test (id, name) VALUES (1, "test");`)

						return err
					},
				},
				wantErr: assert.Error,
			}
		}(),
	}
	for _, tt := range tests {
		t.Run(
			tt.name,
			func(t *testing.T) {
				p := &Postgres{
					db: tt.fields.db,
				}
				tt.wantErr(t, p.DoTransaction(tt.args.ctx, tt.args.fnStmt))
			},
		)
	}
}
