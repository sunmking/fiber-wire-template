package repository

import (
	"fiber-wire-template/pkg/gredis"
	"fiber-wire-template/pkg/log"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type Repository struct {
	db     *dbx.DB
	rdb    *gredis.Redis
	logger *log.Logger
}

func NewRepository(db *dbx.DB, rdb *gredis.Redis, logger *log.Logger) *Repository {
	return &Repository{
		db:     db,
		rdb:    rdb,
		logger: logger,
	}
}
