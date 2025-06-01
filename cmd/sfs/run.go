package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"

	"log/slog"
	"ytils.dev/cli"

	"mkuznets.com/go/sfs/internal/api"
	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/feedstore"
	"mkuznets.com/go/sfs/internal/filestore"
	"mkuznets.com/go/sfs/internal/rss"

	_ "github.com/mattn/go-sqlite3"
)

var _ cli.Commander = (*RunCommand)(nil)

type RunCommand struct {
	DB      *DB      `group:"Database" namespace:"db" env-namespace:"DB"`
	Server  *Server  `group:"Service" namespace:"server" env-namespace:"SERVER" json:"SERVER"`
	Storage *Storage `group:"File storage" namespace:"storage" env-namespace:"STORAGE" json:"STORAGE"`
	Auth    *Auth    `group:"Authentication" namespace:"auth" env-namespace:"AUTH" json:"AUTH"`
}

type DB struct {
	Dsn string `long:"dsn" env:"DSN" description:"Database DSN" required:"true"`
}

type Server struct {
	Addr      string `long:"addr" env:"ADDR" json:"ADDR" description:"HTTP service address" required:"true"`
	UrlPrefix string `long:"url-prefix" env:"URL_PREFIX" json:"URL_PREFIX" description:"URL prefix to the service" required:"true"`
}

type Storage struct {
	S3Opts *S3 `group:"S3" namespace:"s3" env-namespace:"S3" json:"S3"`
}

type S3 struct {
	EndpointUrl  string `long:"endpoint-url" env:"ENDPOINT_URL" description:"endpoint url" json:"ENDPOINT_URL" required:"true"`
	KeyID        string `long:"access-key-id" env:"ACCESS_KEY_ID" description:"access id" json:"ACCESS_KEY_ID" required:"true"`
	SecretKey    string `long:"secret-access-key" env:"SECRET_ACCESS_KEY" description:"access secret" json:"SECRET_ACCESS_KEY" required:"true"`
	Bucket       string `long:"bucket" env:"BUCKET" description:"S3 bucket name" json:"BUCKET" required:"true"`
	UrlTemplate  string `long:"url-template" env:"URL_TEMPLATE" description:"Template of a public URL of the uploaded object" json:"URL_TEMPLATE" required:"true"`
	UsePathStyle bool   `long:"use-path-style" env:"USE_PATH_STYLE" description:"Use path style endpoint URL for S3" json:"USE_PATH_STYLE"`
}

type Auth struct {
	Auth0Opts *Auth0 `group:"Auth0 authentication" namespace:"auth0" env-namespace:"AUTH0" json:"AUTH0"`
}

type Auth0 struct {
	Enabled  bool   `long:"enabled" env:"ENABLED" description:"Enable Auth0 authentication" json:"ENABLED"`
	Domain   string `long:"domain" env:"DOMAIN" description:"Auth0 domain" json:"DOMAIN"`
	Audience string `long:"audience" env:"AUDIENCE" description:"Auth0 audience" json:"AUDIENCE"`
}

func (c *RunCommand) Execute([]string) error {
	var authService auth.Service
	switch {
	case c.Auth.Auth0Opts.Enabled:
		opts := c.Auth.Auth0Opts
		issuerURL, err := url.Parse(fmt.Sprintf("https://%s/", opts.Domain))
		if err != nil {
			return fmt.Errorf("parse auth0 domain: %w", err)
		}
		authService = auth.NewOIDCService(issuerURL, opts.Audience)
	default:
		authService = &auth.NoAuth{}
	}

	fileStorage := filestore.NewS3FileStore(
		c.Storage.S3Opts.EndpointUrl,
		c.Storage.S3Opts.Bucket,
		c.Storage.S3Opts.KeyID,
		c.Storage.S3Opts.SecretKey,
		c.Storage.S3Opts.UrlTemplate,
		c.Storage.S3Opts.UsePathStyle,
	)

	dbDSN, err := prepareDSN(c.DB.Dsn)
	if err != nil {
		return err
	}

	sqlDB, err := sql.Open("sqlite3", dbDSN)
	if err != nil {
		return err
	}

	feedStore := feedstore.NewSQLiteStore(sqlDB)
	if err := feedStore.Init(context.Background()); err != nil {
		return err
	}

	rssController := rss.NewController(feedStore, fileStorage)
	apiController := api.NewController(feedStore, fileStorage, rssController)

	apiService := api.NewService(apiController)

	router := api.NewRouter("/api", authService, apiService)

	slog.Info("starting server", "addr", c.Server.Addr)

	return http.ListenAndServe(c.Server.Addr, router)
}

func prepareDSN(dsn string) (string, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return "", err
	}
	query := u.Query()
	query.Add("_journal_mode", "WAL")
	query.Add("_synchronous", "NORMAL")
	query.Add("_writable_schema", "0")
	query.Add("_foreign_keys", "1")
	u.RawQuery = query.Encode()

	return u.String(), nil
}
