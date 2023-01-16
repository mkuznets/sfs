package main

import (
	"context"
	"database/sql"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"mkuznets.com/go/sps/internal/sps"
	"mkuznets.com/go/sps/internal/sps/api"
	"mkuznets.com/go/sps/internal/sps/feed"
	_ "modernc.org/sqlite"
	"time"
)

type ServerCommand struct {
	server      *sps.Server
	feedService feed.Service
}

func (c *ServerCommand) Init(app *App) error {
	sqldb, err := sql.Open(app.Db.Driver, app.Db.Dsn)
	if err != nil {
		return err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(&hook{})

	store := api.NewStore(db)

	c.server = &sps.Server{
		Addr: ":8080",
		ApiRouter: api.NewRouter(
			api.NewHandler(
				api.NewController(
					store,
					api.NewUploader(),
				),
			),
		),
	}

	c.feedService = feed.NewService(feed.NewController(store))

	return nil
}

func (c *ServerCommand) Execute([]string) error {
	ctx := context.Background()

	go func() {
		c.feedService.Start(ctx)
	}()

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
