package main

import (
	"github.com/jessevdk/go-flags"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sps/internal/sps"
	"os"
)

func main() {
	var opts sps.Flags
	parser := flags.NewParser(&opts, flags.Default)
	parser.CommandHandler = func(command flags.Commander, args []string) error {
		if err := command.Execute(args); err != nil {
			log.Fatal().Msg(err.Error())
			return nil
		}
		return nil
	}

	if _, err := parser.Parse(); err != nil {
		if e, ok := err.(*flags.Error); ok {
			if e.Type == flags.ErrHelp {
				return
			}
		}
		os.Exit(1)
	}

}
