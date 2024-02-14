package feedstore

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"log/slog"
	"ytils.dev/sqlite-migrator"

	"mkuznets.com/go/sfs/sql/sqlite"
)

type contextKey int

var ctxTxKey = contextKey(0x54)

var ErrNotFound = errors.New("not found")

// SQLiteStore implements the FeedStore interface.
type SQLiteStore struct {
	db *bun.DB
}

func NewSQLiteStore(db *sql.DB) *SQLiteStore {
	bunDB := bun.NewDB(db, sqlitedialect.New())
	return &SQLiteStore{
		db: bunDB,
	}
}

func (s *SQLiteStore) ctxDb(ctx context.Context) bun.IDB {
	tx := ctx.Value(ctxTxKey)
	if tx != nil {
		return tx.(bun.Tx)
	}
	return s.db
}

func (s *SQLiteStore) Init(ctx context.Context) error {
	m := migrator.New(s.db.DB, sqlite.FS)
	m.WithLogFunc(func(msg string, args ...interface{}) {
		slog.Info(msg, args...)
	})
	return m.Migrate(ctx)
}

func (s *SQLiteStore) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	return s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return fn(context.WithValue(ctx, ctxTxKey, tx))
	})
}

func (s *SQLiteStore) GetFeeds(ctx context.Context, filter *FeedFilter) ([]*Feed, error) {
	feeds := make([]*Feed, 0)
	q := s.ctxDb(ctx).NewSelect().Model(&feeds)

	if len(filter.Ids) > 0 {
		q = q.Where("id IN (?)", bun.In(filter.Ids))
	}
	if len(filter.UserIds) > 0 {
		q = q.Where("user_id IN (?)", bun.In(filter.UserIds))
	}

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return feeds, nil
}

func (s *SQLiteStore) CreateFeeds(ctx context.Context, feeds []*Feed) error {
	if len(feeds) == 0 {
		return errors.New("no feeds to create")
	}
	_, err := s.ctxDb(ctx).NewInsert().Model(&feeds).Returning("id").Exec(ctx)
	return err
}

func (s *SQLiteStore) GetItems(ctx context.Context, filter *ItemFilter) ([]*Item, error) {
	items := make([]*Item, 0)

	q := s.ctxDb(ctx).NewSelect().Model(&items).Relation("File")

	if len(filter.Ids) > 0 {
		q = q.Where("it.id IN (?)", bun.In(filter.Ids))
	}
	if len(filter.FeedIds) > 0 {
		q = q.Where("it.feed_id IN (?)", bun.In(filter.FeedIds))
	}
	if len(filter.UserIds) > 0 {
		q = q.Where("it.user_id IN (?)", bun.In(filter.UserIds))
	}
	q = q.Order("it.id DESC")

	if err := q.Scan(ctx); err != nil {
		return nil, err
	}

	return items, nil
}

func (s *SQLiteStore) CreateItems(ctx context.Context, items []*Item) error {
	if len(items) == 0 {
		return errors.New("no items to create")
	}
	_, err := s.ctxDb(ctx).NewInsert().Model(&items).Returning("id").Exec(ctx)
	return err
}

func (s *SQLiteStore) CreateFile(ctx context.Context, file *File) error {
	_, err := s.ctxDb(ctx).NewInsert().Model(file).Returning("id").Exec(ctx)
	return err
}

func (s *SQLiteStore) UpdateFeeds(ctx context.Context, feeds []*Feed, fields ...string) error {
	if len(feeds) == 0 {
		return errors.New("no feeds to update")
	}

	values := s.ctxDb(ctx).NewValues(&feeds)

	q := s.ctxDb(ctx).NewUpdate().
		With("_data", values).
		Model((*Feed)(nil)).
		TableExpr("_data")

	for _, attr := range fields {
		q = q.Set(fmt.Sprintf("%s = _data.%s", attr, attr))
	}
	q = q.Where("fe.id = _data.id")

	_, err := q.Exec(ctx)
	return err
}

func (s *SQLiteStore) UpdateFiles(ctx context.Context, files []*File, fields ...string) error {
	if len(files) == 0 {
		return errors.New("no files to update")
	}

	values := s.ctxDb(ctx).NewValues(&files)

	q := s.ctxDb(ctx).NewUpdate().
		With("_data", values).
		Model((*File)(nil)).
		TableExpr("_data")

	for _, attr := range fields {
		q = q.Set(fmt.Sprintf("%s = _data.%s", attr, attr))
	}
	q = q.Where("fi.id = _data.id")

	_, err := q.Exec(ctx)
	return err
}

func (s *SQLiteStore) GetFileById(ctx context.Context, id string) (*File, error) {
	var file File
	err := s.ctxDb(ctx).NewSelect().Model(&file).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &file, nil
}
