package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"log/slog"
	"ytils.dev/cli"

	"mkuznets.com/go/sfs/internal/api"
	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/auth/auth0"
	"mkuznets.com/go/sfs/internal/files"
	"mkuznets.com/go/sfs/internal/rss"
	"mkuznets.com/go/sfs/internal/store"
)

var (
	_ cli.Commander   = (*RunCommand)(nil)
	_ cli.Validator   = (*RunCommand)(nil)
	_ cli.Initer[App] = (*RunCommand)(nil)
)

type RunCommand struct {
	ServerOpts  *Server  `group:"Server options" namespace:"server" env-namespace:"SERVER" json:"SERVER"`
	StorageOpts *Storage `group:"File storage" namespace:"storage" env-namespace:"STORAGE" json:"STORAGE"`
	AuthOpts    *Auth    `group:"Authentication options" namespace:"auth" env-namespace:"AUTH" json:"AUTH"`

	api api.Api
}

func (c *RunCommand) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.ServerOpts),
		validation.Field(&c.StorageOpts),
		validation.Field(&c.AuthOpts),
	)
}

type Server struct {
	Addr      string `long:"addr" env:"ADDR" json:"ADDR" description:"HTTP service address" required:"true"`
	UrlPrefix string `long:"url-prefix" env:"URL_PREFIX" json:"URL_PREFIX" description:"URL prefix to the service" required:"true"`
}

func (s *Server) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.UrlPrefix, validation.Required, is.URL),
	)
}

type Storage struct {
	S3Opts *S3 `group:"S3" namespace:"s3" env-namespace:"S3" json:"S3"`
}

func (s *Storage) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.S3Opts),
	)
}

type S3 struct {
	EndpointUrl string `long:"endpoint-url" env:"ENDPOINT_URL" description:"endpoint url" json:"ENDPOINT_URL" required:"true"`
	KeyID       string `long:"access-key-id" env:"ACCESS_KEY_ID" description:"access id" json:"ACCESS_KEY_ID" required:"true"`
	SecretKey   string `long:"secret-access-key" env:"SECRET_ACCESS_KEY" description:"access secret" json:"SECRET_ACCESS_KEY" required:"true"`
	Bucket      string `long:"bucket" env:"BUCKET" description:"S3 bucket name" json:"BUCKET" required:"true"`
	UrlTemplate string `long:"url-template" env:"URL_TEMPLATE" description:"Template of a public URL of the uploaded object" json:"URL_TEMPLATE" required:"true"`
}

func (s3 *S3) Validate() error {
	return validation.ValidateStruct(
		s3,
		validation.Field(&s3.EndpointUrl, validation.Required, is.URL),
	)
}

type Auth struct {
	Auth0Opts *Auth0 `group:"Auth0 authentication" namespace:"auth0" env-namespace:"AUTH0" json:"AUTH0"`
}

func (a *Auth) Validate() error {
	return validation.ValidateStruct(
		a,
		validation.Field(&a.Auth0Opts),
	)
}

type Auth0 struct {
	Enabled  bool   `long:"enabled" env:"ENABLED" description:"Enable Auth0 authentication" json:"ENABLED"`
	Domain   string `long:"domain" env:"DOMAIN" description:"Auth0 domain" json:"DOMAIN"`
	Audience string `long:"audience" env:"AUDIENCE" description:"Auth0 audience" json:"AUDIENCE"`
}

func (a *Auth0) Validate() error {
	if a.Enabled {
		return validation.ValidateStruct(
			a,
			validation.Field(&a.Domain, validation.Required, is.URL),
			validation.Field(&a.Audience, validation.Required),
		)
	}
	return nil
}

func (c *RunCommand) Init(app *App) error {
	db, err := store.NewBunDb(app.DbOpts.Driver, app.DbOpts.Dsn)
	if err != nil {
		return err
	}

	var authService auth.Service
	switch {
	case c.AuthOpts.Auth0Opts.Enabled:
		opts := c.AuthOpts.Auth0Opts
		issuerURL, err := url.Parse(fmt.Sprintf("https://%s/", opts.Domain))
		if err != nil {
			return fmt.Errorf("parse auth0 domain: %w", err)
		}
		authService = auth0.New(issuerURL, opts.Audience)
	default:
		authService = &auth.NoAuth{}
	}

	fileStorage := files.NewS3Storage(
		c.StorageOpts.S3Opts.EndpointUrl,
		c.StorageOpts.S3Opts.Bucket,
		c.StorageOpts.S3Opts.KeyID,
		c.StorageOpts.S3Opts.SecretKey,
		c.StorageOpts.S3Opts.UrlTemplate,
	)

	bunStore := store.NewBunStore(db)
	if err := bunStore.Init(context.Background()); err != nil {
		return err
	}

	rssController := rss.NewController(bunStore, fileStorage)
	apiController := api.NewController(bunStore, fileStorage, api.NewIdService(), rssController)
	c.api = api.New(authService, api.NewHandler(apiController))

	return nil
}

func (c *RunCommand) Execute([]string) error {
	handler := c.api.Handler("/api")

	slog.Info("starting server", "addr", c.ServerOpts.Addr)

	return http.ListenAndServe(c.ServerOpts.Addr, handler)
}
