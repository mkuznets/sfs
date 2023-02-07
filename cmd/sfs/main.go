package main

import (
	"github.com/joho/godotenv"
	"mkuznets.com/go/ytils/ycli"
	"mkuznets.com/go/ytils/ylog"
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

	DbCmd  *DbCommand  `command:"db" description:"Database migration"`
	RunCmd *RunCommand `command:"run" description:"Start the service"`
}

func init() {
	ylog.Setup()
}

func main() {
	_ = godotenv.Load()
	ycli.Main[App]()
}
