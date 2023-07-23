package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/lib/pq"
	sqlTrace "github.com/signalfx/signalfx-go-tracing/contrib/database/sql"

	"habit-tracker/logger"
)

type Environments struct {
	Host     string `envconfig:"POSTGRES_HOST" required:"true"`
	Port     string `envconfig:"POSTGRES_PORT" required:"true"`
	Name     string `envconfig:"POSTGRES_DATABASE" required:"true"`
	Username string `envconfig:"POSTGRES_USER" required:"true"`
	Password string `envconfig:"POSTGRES_PASSWORD" required:"true"`
}

type Connection struct {
	DbPQTimeout        int
	MaxConnections     int
	MaxIdleConnections int
	Host               string
	Port               string
	DBName             string
	User               string
	Password           string
	ConnMaxLifeTime    time.Duration
}

var postgresConnection Connection

func init() {
	environments := Environments{}
	err := envconfig.Process("", &environments)
	if err != nil {
		logger.New("-").Panicf("[postgres][process][err:%s]", err.Error())
	}

	postgresConnection = Connection{
		DbPQTimeout:        30,
		MaxConnections:     10,
		MaxIdleConnections: 1,
		Host:               environments.Host,
		Port:               environments.Port,
		DBName:             environments.Name,
		User:               environments.Username,
		Password:           environments.Password,
		ConnMaxLifeTime:    30 * time.Minute,
	}
}

type Postgres struct {
	db *sql.DB
}

func NewPostgres() *Postgres {
	log := logger.New("-")

	connectionURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=%d",
		postgresConnection.User,
		postgresConnection.Password,
		postgresConnection.Host,
		postgresConnection.Port,
		postgresConnection.DBName,
		"disable",
		postgresConnection.DbPQTimeout,
	)

	sqlTrace.Register("postgres", &pq.Driver{})

	db, err := sqlTrace.Open("postgres", connectionURL)
	if err != nil {
		log.Panicf("[postgres][new][open][err:%s]", err.Error())
	}

	db.SetMaxIdleConns(postgresConnection.MaxIdleConnections)
	db.SetMaxOpenConns(postgresConnection.MaxConnections)
	db.SetConnMaxLifetime(postgresConnection.ConnMaxLifeTime)

	err = db.Ping()
	if err != nil {
		log.Panicf("[postgres][new][ping][err:%s]", err.Error())
	}

	return &Postgres{
		db: db,
	}
}

func (p *Postgres) QueryContext(ctx context.Context, query string) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query)
}

func (p *Postgres) ExecContext(ctx context.Context, query string) (sql.Result, error) {
	return p.db.ExecContext(ctx, query)
}

func (p *Postgres) DoTransaction(ctx context.Context, fnStmt ExecStmt) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = fnStmt(tx)
	if err != nil {
		rollbackError := tx.Rollback()
		if rollbackError != nil {
			return rollbackError
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
