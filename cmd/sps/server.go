package main

import (
	"context"
	"mkuznets.com/go/sps/internal/files"
	"mkuznets.com/go/sps/internal/sps"
	"mkuznets.com/go/sps/internal/sps/api"
	"mkuznets.com/go/sps/internal/sps/feed"
	"mkuznets.com/go/sps/internal/store"
	_ "modernc.org/sqlite"
)

type ServerCommand struct {
	ServerOpts Server `group:"Server Options" namespace:"server" env-namespace:"SERVER"`
	AwsOpts    Aws    `group:"AWS Options" namespace:"aws" env-namespace:"AWS"`

	server      *sps.Server
	feedService feed.Service
}

type Server struct {
	Addr      string `long:"addr" env:"ADDR" description:"HTTP service address" required:"true"`
	UrlPrefix string `long:"url-prefix" env:"URL_PREFIX" description:"URL prefix to the service" required:"true"`
}

type Aws struct {
	EndpointUrl string `long:"endpoint-url" env:"ENDPOINT_URL" description:"endpoint url" required:"true"`
	KeyID       string `long:"access-key-id" env:"ACCESS_KEY_ID" description:"access id" required:"true"`
	SecretKey   string `long:"secret-access-key" env:"SECRET_ACCESS_KEY" description:"access secret" required:"true"`
	Bucket      string `long:"s3-bucket" env:"S3_BUCKET" description:"S3 bucket name" required:"true"`
	UrlPrefix   string `long:"url-prefix" env:"URL_PREFIX" description:"Optional URL prefix that will be used to generate URLs to the uploaded files"`
}

func (c *ServerCommand) Init(app *App) error {
	db, err := store.NewBunDb(app.DbOpts.Driver, app.DbOpts.Dsn)
	if err != nil {
		return err
	}

	fileStorage := files.NewS3Storage(c.AwsOpts.EndpointUrl, c.AwsOpts.Bucket, c.AwsOpts.KeyID, c.AwsOpts.SecretKey, c.AwsOpts.UrlPrefix)
	bunStore := store.NewBunStore(db)
	feedController := feed.NewController(bunStore, fileStorage)

	c.server = &sps.Server{
		Addr: c.ServerOpts.Addr,
		ApiRouter: api.NewRouter(
			api.NewHandler(
				api.NewController(
					bunStore,
					fileStorage,
					api.NewIdService(),
					feedController,
				),
			),
		),
	}

	c.feedService = feed.NewService(feedController)

	return nil
}

func (c *ServerCommand) Execute([]string) error {
	ctx := context.Background()

	go func() {
		c.feedService.Start(ctx)
	}()

	return c.server.Start()
}
