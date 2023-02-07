package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"golang.org/x/exp/slog"
	"mkuznets.com/go/ytils/ylog"
	"net/url"
	"time"

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
	db.AddQueryHook(&bunHook{})

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

type bunHook struct{}

func (h *bunHook) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context {
	return ctx
}

func (h *bunHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	dur := time.Since(event.StartTime)

	logAttrs := []slog.Attr{
		slog.Duration("duration", dur),
		slog.String("query", event.Query),
	}
	if event.Err != nil {
		logAttrs = append(logAttrs, slog.Any(slog.ErrorKey, event.Err))
	}
	level := slog.LevelInfo

	if event.Err != nil && !errors.Is(event.Err, sql.ErrNoRows) {
		level = slog.LevelError
	}

	ylog.Ctx(ctx).LogAttrs(level, "QUERY", logAttrs...)
}
