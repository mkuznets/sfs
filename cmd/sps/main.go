package main

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/jessevdk/go-flags"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sps/internal/ytils/ylog"
	"os"
)

// Global is a group of common flags for all subcommands.
type Global struct{}

type Db struct {
	Driver string `long:"driver" env:"DRIVER" description:"Database driver" default:"sqlite3" choice:"sqlite3" required:"true"`
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
	flags.Commander
	validation.Validatable
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
	ylog.Setup()
}

func main() {
	validation.ErrorTag = "validation"

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	var app App
	var parser = flags.NewParser(&app, flags.Default)
	parser.CommandHandler = func(command flags.Commander, args []string) error {
		c := command.(Command)
		handleError(c.Validate())
		handleError(c.Init(&app))
		handleError(c.Execute(args))
		return nil
	}

	if _, err := parser.Parse(); err != nil {
		handleError(err)
	}
}
