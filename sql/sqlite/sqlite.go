package sqlite

import (
	"embed"
	"github.com/uptrace/bun/migrate"
)

//go:embed *.sql

var files embed.FS

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(files); err != nil {
		panic(err)
	}
}
