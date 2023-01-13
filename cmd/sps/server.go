package main

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"mkuznets.com/go/sps/internal/sps"
	"mkuznets.com/go/sps/internal/sps/api"
	_ "modernc.org/sqlite"
	"time"
)

type ServerCommand struct {
	server *sps.Server
}

func (c *ServerCommand) Init(app *App) error {
	sqldb, err := sql.Open(app.Db.Driver, app.Db.Dsn)
	if err != nil {
		return err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(&hook{})

	c.server = &sps.Server{
		Addr: ":8080",
		ApiRouter: api.NewRouter(
			api.NewHandler(
				api.NewController(
					api.NewStore(db),
					api.NewUploader(),
				),
			),
		),
	}
	return nil
}

func (c *ServerCommand) Execute([]string) error {
	return c.server.Start()
}

type hook struct {
}

func (h *hook) BeforeQuery(ctx context.Context, _ *bun.QueryEvent) context.Context {
	return ctx
}

func (h *hook) AfterQuery(_ context.Context, event *bun.QueryEvent) {
	dur := time.Since(event.StartTime)
	l := log.Debug().Str("query", event.Query).Dur("duration", dur)
	if event.Err != nil {
		l = l.Err(event.Err)
	}
	l.Msg("bun query")
}
