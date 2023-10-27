package store

import (
	"database/sql"
	"net/url"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"

	// Required to load "sqlite" driver
	_ "github.com/mattn/go-sqlite3"
)

func NewBunDb(driver, dsn string) (*bun.DB, error) {
	dsn, err := prepareDsn(driver, dsn)
	if err != nil {
		return nil, err
	}

	sqldb, err := sql.Open(driver, dsn)
	if err != nil {
		return nil, err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())

	return db, nil
}

func prepareDsn(driver, dsn string) (string, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return "", err
	}
	query := u.Query()

	if driver == "sqlite3" {
		query.Add("_journal_mode", "WAL")
		query.Add("_synchronous", "NORMAL")
		query.Add("_writable_schema", "0")
		query.Add("_foreign_keys", "1")
	}
	u.RawQuery = query.Encode()

	return u.String(), nil
}
