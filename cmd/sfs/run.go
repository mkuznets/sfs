package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"

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
	S3Opts    *S3           `group:"S3" namespace:"s3" env-namespace:"S3" json:"S3"`
	LocalOpts *LocalStorage `group:"Local storage" namespace:"local" env-namespace:"LOCAL" json:"LOCAL"`
}

func (s *Storage) Validate() error {
	if s.S3Opts.Enabled && s.LocalOpts.Enabled {
		return errors.New("only one storage type can be enabled")
	}
	if !s.S3Opts.Enabled && !s.LocalOpts.Enabled {
		return errors.New("at least one storage type must be enabled")
	}

	return validation.ValidateStruct(
		s,
		validation.Field(&s.S3Opts),
		validation.Field(&s.LocalOpts),
	)
}

type LocalStorage struct {
	Enabled bool   `long:"enabled" env:"ENABLED" description:"Enable local storage" json:"ENABLED"`
	Path    string `long:"path" env:"PATH" description:"Path to the local file storage directory" json:"PATH"`
}

func (l *LocalStorage) Validate() error {
	if l.Enabled {
		return validation.ValidateStruct(
			l,
			validation.Field(&l.Path, validation.Required, validation.By(validateDirectory)),
		)
	}
	return nil
}

func validateDirectory(x interface{}) error {
	s := x.(string)
	if err := os.MkdirAll(s, 0o755); err != nil {
		return fmt.Errorf("invalid directory: %w", err)
	}
	return nil
}

type S3 struct {
	Enabled     bool   `long:"enabled" env:"ENABLED" description:"Enable S3 storage" json:"ENABLED"`
	EndpointUrl string `long:"endpoint-url" env:"ENDPOINT_URL" description:"endpoint url" json:"ENDPOINT_URL"`
	KeyID       string `long:"access-key-id" env:"ACCESS_KEY_ID" description:"access id" json:"ACCESS_KEY_ID"`
	SecretKey   string `long:"secret-access-key" env:"SECRET_ACCESS_KEY" description:"access secret" json:"SECRET_ACCESS_KEY"`
	Bucket      string `long:"bucket" env:"BUCKET" description:"S3 bucket name" json:"BUCKET"`
	UrlTemplate string `long:"url-template" env:"URL_TEMPLATE" description:"Template of a publically available URL of the uploaded object" json:"URL_TEMPLATE"`
}

func (s3 *S3) Validate() error {
	if s3.Enabled {
		return validation.ValidateStruct(
			s3,
			validation.Field(&s3.EndpointUrl, validation.Required, is.URL),
			validation.Field(&s3.EndpointUrl, validation.Required),
			validation.Field(&s3.KeyID, validation.Required),
			validation.Field(&s3.SecretKey, validation.Required),
			validation.Field(&s3.Bucket, validation.Required),
			validation.Field(&s3.UrlTemplate, validation.Required),
		)
	}
	return nil
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

	var fileStorage files.Storage
	switch {
	case c.StorageOpts.S3Opts.Enabled:
		fileStorage = files.NewS3Storage(
			c.StorageOpts.S3Opts.EndpointUrl,
			c.StorageOpts.S3Opts.Bucket,
			c.StorageOpts.S3Opts.KeyID,
			c.StorageOpts.S3Opts.SecretKey,
			c.StorageOpts.S3Opts.UrlTemplate,
		)
	case c.StorageOpts.LocalOpts.Enabled:
		fileStorage = files.NewLocalStorage(c.StorageOpts.LocalOpts.Path, c.ServerOpts.UrlPrefix)
	}

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
