package store

import (
	"context"
	"github.com/uptrace/bun"
	"mkuznets.com/go/sps/internal/ytils/ytime"
)

type Store interface {
	Init(ctx context.Context) error

	GetFeeds(ctx context.Context, filter *FeedFilter) ([]*Feed, error)
	CreateFeeds(ctx context.Context, feeds []*Feed) error
	UpdateFeeds(ctx context.Context, feeds []*Feed, fields ...string) error

	GetItems(ctx context.Context, filter *ItemFilter) ([]*Item, error)
	CreateItems(ctx context.Context, items []*Item) error

	GetFileById(ctx context.Context, id string) (*File, error)
	CreateFile(ctx context.Context, file *File) error
	UpdateFiles(ctx context.Context, files []*File, fields ...string) error

	Tx(ctx context.Context, fn func(ctx context.Context) error) error
}

type Feed struct {
	bun.BaseModel `bun:"table:feeds,alias:fe"`

	Id           string     `bun:"id,pk"`
	UserId       string     `bun:"user_id"`
	Type         string     `bun:"type"`
	Title        string     `bun:"title"`
	Link         string     `bun:"link"`
	Authors      string     `bun:"authors"`
	Description  string     `bun:"description"`
	Rss          string     `bun:"rss"`
	RssUpdatedAt ytime.Time `bun:"rss_updated_at"`
	CreatedAt    ytime.Time `bun:"created_at"`
	UpdatedAt    ytime.Time `bun:"updated_at"`
	DeletedAt    ytime.Time `bun:"deleted_at"`
}

type Item struct {
	bun.BaseModel `bun:"table:items,alias:it"`

	Id          string     `bun:"id,pk"`
	FeedId      string     `bun:"feed_id"`
	UserId      string     `bun:"user_id"`
	FileId      string     `bun:"file_id"`
	File        *File      `bun:"rel:belongs-to,join:file_id=id"`
	Title       string     `bun:"title"`
	Description string     `bun:"description"`
	Link        string     `bun:"link"`
	Authors     string     `bun:"authors"`
	CreatedAt   ytime.Time `bun:"created_at"`
	UpdatedAt   ytime.Time `bun:"updated_at"`
	DeletedAt   ytime.Time `bun:"deleted_at"`
}

type File struct {
	bun.BaseModel `bun:"table:files,alias:fi"`

	Id        string     `bun:"id,pk"`
	UserId    string     `bun:"user_id"`
	ItemId    string     `bun:"item_id"`
	Size      int64      `bun:"size"`
	MimeType  string     `bun:"mime_type"`
	Hash      string     `bun:"hash"`
	UploadUrl string     `bun:"upload_url"`
	UploadId  string     `bun:"upload_id"`
	CreatedAt ytime.Time `bun:"created_at"`
	UpdatedAt ytime.Time `bun:"updated_at"`
	DeletedAt ytime.Time `bun:"deleted_at"`
}

type FeedFilter struct {
	Ids     []string
	UserIds []string
}

type ItemFilter struct {
	Ids     []string
	FeedIds []string
	UserIds []string
}
