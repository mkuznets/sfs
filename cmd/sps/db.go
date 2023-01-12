package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/migrate"
	"mkuznets.com/go/sps/internal/migrator"
	ssql "mkuznets.com/go/sps/sql"
	_ "modernc.org/sqlite"
)

type DbCommand struct {
	m migrator.Migrator
}

func (c *DbCommand) Init(app *App) error {
	sqldb, err := sql.Open(app.Db.Driver, app.Db.Dsn)
	if err != nil {
		return err
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	c.m = migrator.New(migrate.NewMigrator(db, ssql.Migrations, migrate.WithMarkAppliedOnSuccess(true)))
	return nil
}

func (c *DbCommand) Usage() string {
	return "OP\nWhere OP is one of: init, up, down, lock, unlock, create, status, mark-applied"
}

func (c *DbCommand) Execute(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing OP argument")
	}

	ctx := context.Background()
	op := args[0]

	switch op {
	case "init":
		return c.m.Init(ctx)
	case "up":
		return c.m.Migrate(ctx)
	case "down":
		return c.m.Rollback(ctx)
	case "lock":
		return c.m.Lock(ctx)
	case "unlock":
		return c.m.Unlock(ctx)
	case "create":
		if len(args) < 2 {
			return fmt.Errorf("name argument is required")
		}
		return c.m.CreateSQLMigrations(ctx, args[1])
	case "status":
		return c.m.Status(ctx)
	case "mark-applied":
		return c.m.MarkApplied(ctx)
	default:
		return fmt.Errorf("unknown operation: %s", op)
	}
}
