package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

// Global is a group of common flags for all subcommands.
type Global struct {
	Debug bool `long:"debug" description:"Enable debug logging"`
}

type Db struct {
	Driver string `long:"driver" env:"DRIVER" description:"Database driver" default:"sqlite" required:"true"`
	Dsn    string `long:"dsn" env:"DSN" description:"Database DSN" required:"true"`
}

type App struct {
	GlobalOpts *Global `group:"Global Options"`
	DbOpts     *Db     `group:"Database Options" namespace:"db" env-namespace:"DB"`

	DbCmd     *DbCommand     `command:"db" description:"Database migration"`
	ServerCmd *ServerCommand `command:"server" description:"Start the server"`
}

type Command interface {
	Init(app *App) error
	Execute(args []string) error
}

func handleError(err error) {
	if err == nil {
		return
	}
	if e, ok := err.(*flags.Error); ok {
		if e.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	log.Fatal().Err(err).Msg("fatal error")
	return
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	})
	zerolog.DefaultContextLogger = &log.Logger
	zerolog.DurationFieldUnit = time.Microsecond
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var app App
	var parser = flags.NewParser(&app, flags.Default)
	parser.CommandHandler = func(command flags.Commander, args []string) error {
		c := command.(Command)
		handleError(c.Init(&app))
		handleError(c.Execute(args))
		return nil
	}

	if _, err := parser.Parse(); err != nil {
		handleError(err)
	}
}
