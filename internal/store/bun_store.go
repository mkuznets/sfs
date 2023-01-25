package store

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/uptrace/bun"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
)

type contextKey int

var ctxTxKey = contextKey(0x54)

// bunStore implements the Store interface.
type bunStore struct {
	db *bun.DB
}

func NewBunStore(db *bun.DB) Store {
	return &bunStore{
		db: db,
	}
}

func (s *bunStore) ctxDb(ctx context.Context) bun.IDB {
	tx := ctx.Value(ctxTxKey)
	if tx != nil {
		return tx.(bun.Tx)
	}
	return s.db
}

func (s *bunStore) Init(ctx context.Context) error {
	m := NewBunMigrator(s.db)
	if err := m.Init(ctx); err != nil {
		return err
	}
	if err := m.Migrate(ctx); err != nil {
		return err
	}
	return nil
}

func (s *bunStore) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	return s.db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		return fn(context.WithValue(ctx, ctxTxKey, tx))
	})
}

func (s *bunStore) GetFeeds(ctx context.Context, filter *FeedFilter) ([]*Feed, error) {
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

func (s *bunStore) CreateFeeds(ctx context.Context, feeds []*Feed) error {
	if len(feeds) == 0 {
		return yerr.New("no feeds to create")
	}
	_, err := s.ctxDb(ctx).NewInsert().Model(&feeds).Returning("id").Exec(ctx)
	return err
}

func (s *bunStore) GetItems(ctx context.Context, filter *ItemFilter) ([]*Item, error) {
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

func (s *bunStore) CreateItems(ctx context.Context, items []*Item) error {
	if len(items) == 0 {
		return yerr.New("no items to create")
	}
	_, err := s.ctxDb(ctx).NewInsert().Model(&items).Returning("id").Exec(ctx)
	return err
}

func (s *bunStore) CreateFile(ctx context.Context, file *File) error {
	_, err := s.ctxDb(ctx).NewInsert().Model(file).Returning("id").Exec(ctx)
	return err
}

func (s *bunStore) UpdateFeeds(ctx context.Context, feeds []*Feed, fields ...string) error {
	if len(feeds) == 0 {
		return yerr.New("no feeds to update")
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

func (s *bunStore) UpdateFiles(ctx context.Context, files []*File, fields ...string) error {
	if len(files) == 0 {
		return yerr.New("no files to update")
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

func (s *bunStore) GetFileById(ctx context.Context, id string) (*File, error) {
	var file File
	err := s.ctxDb(ctx).NewSelect().Model(&file).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, yerr.NotFound("file not found")
		}
		return nil, err
	}
	return &file, nil
}
