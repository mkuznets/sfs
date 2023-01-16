package main

import (
	"github.com/jessevdk/go-flags"
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
	Driver string `long:"driver" description:"Database driver" default:"sqlite"`
	Dsn    string `long:"dsn" description:"Database Dsn" default:"file:./data/sps.db?cache=shared&mode=rwc&_pragma=journal_mode(WAL)"`
}

type App struct {
	Global *Global        `group:"Global Options"`
	Db     *Db            `group:"Database Options"`
	Dbm    *DbCommand     `command:"db" description:"Database migration"`
	Server *ServerCommand `command:"server" description:"Start the server"`
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

func setupLogger() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "2006-01-02 15:04:05",
	})
	zerolog.DefaultContextLogger = &log.Logger
	zerolog.DurationFieldUnit = time.Microsecond
}

func main() {
	var app App
	var parser = flags.NewParser(&app, flags.Default)
	parser.CommandHandler = func(command flags.Commander, args []string) error {
		setupLogger()
		c := command.(Command)
		handleError(c.Init(&app))
		handleError(c.Execute(args))
		return nil
	}

	if _, err := parser.Parse(); err != nil {
		handleError(err)
	}
}
