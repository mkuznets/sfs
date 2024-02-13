package main

import (
	"github.com/joho/godotenv"
	"mkuznets.com/go/ytils/ycli"

	"mkuznets.com/go/sfs/internal/slogger"
)

type Db struct {
	Driver string `long:"driver" env:"DRIVER" description:"Database driver" default:"sqlite3" choice:"sqlite3" required:"true"`
	Dsn    string `long:"dsn" env:"DSN" description:"Database DSN" required:"true"`
}

type App struct {
	DbOpts *Db `group:"Database Options" namespace:"db" env-namespace:"DB"`

	RunCmd *RunCommand `command:"run" description:"Start the service"`
}

func main() {
	_ = godotenv.Load()
	slogger.Init()
	ycli.Main[App]()
}
