package ozzodb

import (
	"context"
	"database/sql"
	"fiber-wire-template/pkg/config"
	"fiber-wire-template/pkg/log"
	"fmt"
	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"os"
	"strconv"
	"time"
)

// DB represents a DB connection that can be used to run SQL queries.
type DB struct {
	db *dbx.DB
}

var sqlDb *dbx.DB

// TransactionFunc represents a function that will start a transaction and run the given function.
type TransactionFunc func(ctx context.Context, f func(ctx context.Context) error) error

type contextKey int

const (
	txKey contextKey = iota
)

func NewDb(conf *config.Config, logger *log.Logger) *dbx.DB {
	var err error
	// MySQL 数据库地址和用户信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DatabaseCnf.User,
		conf.DatabaseCnf.Password,
		conf.DatabaseCnf.Host,
		conf.DatabaseCnf.Port,
		conf.DatabaseCnf.DbName,
	)
	// connect to the database
	db, err := dbx.MustOpen("mysql", dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(-1)
	}
	db.QueryLogFunc = logDBQuery(logger)
	db.ExecLogFunc = logDBExec(logger)

	sqlDb = db

	return db
}
func CloseDb() {
	err := sqlDb.Close()
	if err != nil {
		return
	}
}

// New returns a new DB connection that wraps the given dbx.DB instance.
func New(db *dbx.DB) *DB {
	return &DB{db}
}

// DB returns the dbx.DB wrapped by this object.
func (db *DB) DB() *dbx.DB {
	return db.db
}

// With returns a Builder that can be used to build and execute SQL queries.
// With will return the transaction if it is found in the given context.
// Otherwise, it will return a DB connection associated with the context.
func (db *DB) With(ctx context.Context) dbx.Builder {
	if tx, ok := ctx.Value(txKey).(*dbx.Tx); ok {
		return tx
	}
	return db.db.WithContext(ctx)
}

// Transactional starts a transaction and calls the given function with a context storing the transaction.
// The transaction associated with the context can be access via With().
func (db *DB) Transactional(ctx context.Context, f func(ctx context.Context) error) error {
	return db.db.TransactionalContext(ctx, nil, func(tx *dbx.Tx) error {
		return f(context.WithValue(ctx, txKey, tx))
	})
}

// logDBQuery returns a logging function that can be used to log SQL queries.
func logDBQuery(logger *log.Logger) dbx.QueryLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, rows *sql.Rows, err error) {
		if err == nil {
			logger.With(ctx, zap.String("duration", strconv.FormatInt(t.Milliseconds(), 10)), zap.String("sql", sql)).Info("DB query successful")
		} else {
			logger.With(ctx, zap.String("sql", sql)).Error("DB query error: ${err}")
		}
	}
}

// logDBExec returns a logging function that can be used to log SQL executions.
func logDBExec(logger *log.Logger) dbx.ExecLogFunc {
	return func(ctx context.Context, t time.Duration, sql string, result sql.Result, err error) {
		if err == nil {
			logger.With(ctx, zap.String("duration", strconv.FormatInt(t.Milliseconds(), 10)), zap.String("sql", sql)).Info("DB execution successful")
		} else {
			logger.With(ctx, zap.String("sql", sql)).Error("DB execution error: %err")
		}
	}
}
