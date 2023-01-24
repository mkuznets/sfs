package main

import (
	"context"
	"fmt"
	"mkuznets.com/go/sfs/internal/store"
)

type DbCommand struct {
	m store.Migrator
}

func (c *DbCommand) Init(app *App) error {
	db, err := store.NewBunDb(app.DbOpts.Driver, app.DbOpts.Dsn)
	if err != nil {
		return err
	}
	c.m = store.NewBunMigrator(db)
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
