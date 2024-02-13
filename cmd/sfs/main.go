package main

import (
	"github.com/joho/godotenv"
	"ytils.dev/cli"

	"mkuznets.com/go/sfs/internal/slogger"
)

type App struct {
	Run *RunCommand `command:"run" description:"Start the service"`
}

func main() {
	_ = godotenv.Load()
	slogger.Init()
	cli.ParseExecute[App]()
}
