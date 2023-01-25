package main

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"mkuznets.com/go/sfs/internal/api"
	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/files"
	"mkuznets.com/go/sfs/internal/rss"
	"mkuznets.com/go/sfs/internal/store"
	"mkuznets.com/go/sfs/internal/ytils/ycrypto"
	"net/http"
)

type RunCommand struct {
	ServerOpts *Server `group:"Server Options" namespace:"server" env-namespace:"SERVER" json:"SERVER"`
	S3Opts     *S3     `group:"S3 Options" namespace:"s3" env-namespace:"S3" json:"S3"`
	JwtOpts    *Jwt    `group:"JWT Options" namespace:"jwt" env-namespace:"JWT" json:"JWT"`

	api api.Api
}

func (c *RunCommand) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.ServerOpts),
		validation.Field(&c.S3Opts),
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
			validation.Field(&s3.UrlTemplate, validation.Required, is.URL),
		)
	}
	return nil
}

type Jwt struct {
	PublicKey  string `long:"public-key" env:"PUBLIC_KEY" description:"RSA public key" required:"true"`
	PrivateKey string `long:"private-key" env:"PRIVATE_KEY" description:"RSA private key" required:"true"`
}

func (c *RunCommand) Init(app *App) error {
	db, err := store.NewBunDb(app.DbOpts.Driver, app.DbOpts.Dsn)
	if err != nil {
		return err
	}

	authService := auth.New(c.JwtOpts.PrivateKey, c.JwtOpts.PublicKey)

	fileStorage := files.NewS3Storage(
		c.S3Opts.EndpointUrl,
		c.S3Opts.Bucket,
		ycrypto.MustReveal(c.S3Opts.KeyID),
		ycrypto.MustReveal(c.S3Opts.SecretKey),
		c.S3Opts.UrlTemplate,
	)
	bunStore := store.NewBunStore(db)
	if err := bunStore.Init(context.Background()); err != nil {
		return err
	}

	feedController := rss.NewController(bunStore, fileStorage)
	apiController := api.NewController(bunStore, fileStorage, api.NewIdService(), feedController, authService)
	c.api = api.New(authService, api.NewHandler(apiController))

	return nil
}

func (c *RunCommand) Execute([]string) error {
	handler := c.api.Handler("/api")
	return http.ListenAndServe(c.ServerOpts.Addr, handler)
}
