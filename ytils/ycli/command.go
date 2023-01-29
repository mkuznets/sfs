package ycli

import (
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"
	"os"
)

type Command[T any] interface {
	Init(app *T) error
	Validate() error
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
	log.Fatal().Err(err).Send()
	return
}

func Execute[T any]() {
	log.Debug().Int("pid", os.Getpid()).Msgf("%s started", os.Args[0])

	var app T
	var parser = flags.NewParser(&app, flags.Default)
	parser.CommandHandler = func(command flags.Commander, args []string) error {
		c := command.(Command[T])
		handleError(c.Validate())
		handleError(c.Init(&app))
		handleError(c.Execute(args))
		return nil
	}

	if _, err := parser.Parse(); err != nil {
		handleError(err)
	}
}
