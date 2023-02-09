package store

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"golang.org/x/exp/slog"
	"mkuznets.com/go/sfs/sql/sqlite"
	"os"
	"regexp"
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

var migrationNameRegex = regexp.MustCompile(`(up|down)\.sql$`)

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
		slog.Info("database is up to date")
		return nil
	}
	slog.Info("migrated", "group", group)
	return nil
}

func (m *migrator) Rollback(ctx context.Context) error {
	group, err := m.bm.Rollback(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		slog.Info("there are no groups to roll back")
		return nil
	}
	slog.Info("rolled back", "group", group)
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

	// Make sure all migrations are transactional.
	for _, file := range files {
		newPath := migrationNameRegex.ReplaceAllString(file.Path, "tx.$1.sql")
		if err := os.Rename(file.Path, newPath); err != nil {
			return fmt.Errorf("rename %s to %s: %w", file.Path, newPath, err)
		}
		file.Path = newPath
	}

	for _, mf := range files {
		slog.Info("created migration", "name", mf.Name, "path", mf.Path)
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
		slog.Info("no new migrations to mark as applied")
		return nil
	}
	slog.Info("marked as applied", "group", group)
	return nil
}
