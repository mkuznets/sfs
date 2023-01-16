package main

import (
	"context"
	"mkuznets.com/go/sps/internal/sps"
	"mkuznets.com/go/sps/internal/sps/api"
	"mkuznets.com/go/sps/internal/sps/feed"
	"mkuznets.com/go/sps/internal/store"
	_ "modernc.org/sqlite"
)

type ServerCommand struct {
	server      *sps.Server
	feedService feed.Service
}

func (c *ServerCommand) Init(app *App) error {
	db, err := store.NewBunDb(app.Db.Driver, app.Db.Dsn)
	if err != nil {
		return err
	}

	bunStore := store.NewBunStore(db)

	c.server = &sps.Server{
		Addr: ":8080",
		ApiRouter: api.NewRouter(
			api.NewHandler(
				api.NewController(
					bunStore,
					api.NewUploader(),
					api.NewIdService(),
				),
			),
		),
	}

	c.feedService = feed.NewService(feed.NewController(bunStore))

	return nil
}

func (c *ServerCommand) Execute([]string) error {
	ctx := context.Background()

	go func() {
		c.feedService.Start(ctx)
	}()

	return c.server.Start()
}
