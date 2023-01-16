package store

import (
	"context"
	"github.com/uptrace/bun"
	"mkuznets.com/go/sps/internal/types"
)

type Store interface {
	CreateChannel(ctx context.Context, channel *Channel) error
	GetChannel(ctx context.Context, id string) (*Channel, error)
	ListChannels(ctx context.Context, userId string) ([]*Channel, error)
	GetChannelsIdsToUpdateFeeds(ctx context.Context) ([]string, error)
	UpdateChannelFeeds(ctx context.Context, channels []*Channel) error

	CreateEpisode(ctx context.Context, episode *Episode) error
	ListEpisodesWithFiles(ctx context.Context, channelId string) ([]*Episode, error)

	CreateFile(ctx context.Context, file *File) error
	GetFile(ctx context.Context, id string) (*File, error)
}

type Channel struct {
	bun.BaseModel `bun:"table:channels,alias:ch"`

	Id          string     `bun:"id,pk"`
	UserId      string     `bun:"user_id"`
	Title       string     `bun:"title"`
	Link        string     `bun:"link"`
	Authors     string     `bun:"authors"`
	Description string     `bun:"description"`
	CreatedAt   types.Time `bun:"created_at"`
	UpdatedAt   types.Time `bun:"updated_at"`
	DeletedAt   types.Time `bun:"deleted_at"`
	Feed        Feed       `bun:"embed:feed_"`
}

type Feed struct {
	Content     []byte     `bun:"content"`
	Url         string     `bun:"url"`
	PublishedAt types.Time `bun:"published_at"`
}

type Episode struct {
	bun.BaseModel `bun:"table:episodes,alias:ep"`

	Id          string      `bun:"id,pk"`
	ChannelId   string      `bun:"channel_id"`
	Title       string      `bun:"title"`
	Description string      `bun:"description"`
	Link        string      `bun:"link"`
	Authors     string      `bun:"authors"`
	FileId      string      `bun:"file_id"`
	File        *File       `bun:"rel:belongs-to,join:file_id=id"`
	CreatedAt   types.Time  `bun:"created_at"`
	UpdatedAt   types.Time  `bun:"updated_at"`
	DeletedAt   *types.Time `bun:"deleted_at"`
}

type File struct {
	bun.BaseModel `bun:"table:files,alias:f"`

	Id          string      `bun:"id,pk"`
	UserId      string      `bun:"user_id"`
	Url         string      `bun:"url"`
	Size        int64       `bun:"size"`
	ContentType string      `bun:"content_type"`
	CreatedAt   types.Time  `bun:"created_at"`
	UpdatedAt   types.Time  `bun:"updated_at"`
	DeletedAt   *types.Time `bun:"deleted_at"`
}
