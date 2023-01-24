package store

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"mkuznets.com/go/sfs/sql/sqlite"
)

type Migrator interface {
	Init(ctx context.Context) error
	Migrate(ctx context.Context) error
	Rollback(ctx context.Context) error
	Lock(ctx context.Context) error
	Unlock(ctx context.Context) error
	CreateSQLMigrations(ctx context.Context, name string) error
	Status(ctx context.Context) error
	MarkApplied(ctx context.Context) error
}

type migrator struct {
	bm *migrate.Migrator
}

func NewBunMigrator(db *bun.DB) Migrator {
	return &migrator{
		bm: migrate.NewMigrator(db, sqlite.Migrations, migrate.WithMarkAppliedOnSuccess(true)),
	}
}

func (m *migrator) Init(ctx context.Context) error {
	return m.bm.Init(ctx)
}

func (m *migrator) Migrate(ctx context.Context) error {
	group, err := m.bm.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return nil
	}
	fmt.Printf("migrated to %s\n", group)
	return nil
}

func (m *migrator) Rollback(ctx context.Context) error {
	group, err := m.bm.Rollback(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no groups to roll back\n")
		return nil
	}
	fmt.Printf("rolled back %s\n", group)
	return nil
}

func (m *migrator) Lock(ctx context.Context) error {
	return m.bm.Lock(ctx)
}

func (m *migrator) Unlock(ctx context.Context) error {
	return m.bm.Unlock(ctx)
}

func (m *migrator) CreateSQLMigrations(ctx context.Context, name string) error {
	files, err := m.bm.CreateSQLMigrations(ctx, name)
	if err != nil {
		return err
	}

	for _, mf := range files {
		fmt.Printf("created migration %s (%s)\n", mf.Name, mf.Path)
	}
	return nil
}

func (m *migrator) Status(ctx context.Context) error {
	ms, err := m.bm.MigrationsWithStatus(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("migrations: %s\n", ms)
	fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
	fmt.Printf("last migration group: %s\n", ms.LastGroup())
	return nil
}

func (m *migrator) MarkApplied(ctx context.Context) error {
	group, err := m.bm.Migrate(ctx, migrate.WithNopMigration())
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to mark as applied\n")
		return nil
	}
	fmt.Printf("marked as applied %s\n", group)
	return nil
}
