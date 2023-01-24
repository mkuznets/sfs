package store

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rs/zerolog"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
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

	if driver == "sqlite" {
		query.Add("_journal_mode", "WAL")
		query.Add("_synchronous", "NORMAL")
		query.Add("_writable_schema", "0")
		query.Add("_foreign_keys", "1")

		// Parameters for modernc.org/sqlite/lib
		//query.Add("_pragma", "journal_mode('WAL')")
		//query.Add("_pragma", "synchronous('NORMAL')")
		//query.Add("_pragma", "writable_schema('OFF')")
		//query.Add("_pragma", "encoding('UTF-8')")
		//query.Add("_pragma", "foreign_keys(1)")
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
	var logEvent *zerolog.Event
	l := zerolog.Ctx(ctx)

	if event.Err != nil {
		if errors.Is(event.Err, sql.ErrNoRows) {
			logEvent = l.Warn().Err(event.Err)
		} else {
			logEvent = l.Err(event.Err)
		}
	} else {
		logEvent = l.Debug()
	}

	logEvent.Dur("duration", dur).Str("query", event.Query).Msg("query")
}
