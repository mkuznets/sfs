package store

import (
	"context"
	"github.com/uptrace/bun"
	"mkuznets.com/go/sps/internal/ytils/ytime"
)

type Store interface {
	CreateChannel(ctx context.Context, channel *Channel) error
	GetChannel(ctx context.Context, id string) (*Channel, error)
	ListChannels(ctx context.Context, userId string) ([]*Channel, error)
	GetChannelsIdsToUpdateFeeds(ctx context.Context) ([]string, error)
	UpdateChannelFeeds(ctx context.Context, channels []*Channel) error

	CreateEpisode(ctx context.Context, episode *Episode) error
	GetEpisode(ctx context.Context, id string) (*Episode, error)
	ListEpisodesWithFiles(ctx context.Context, channelId string) ([]*Episode, error)

	CreateFile(ctx context.Context, file *File) error
	GetFile(ctx context.Context, id string) (*File, error)

	CreateUser(ctx context.Context, user *User) error
	GetUserByAccountNumber(ctx context.Context, accountNumber string) (*User, error)
}

type Channel struct {
	bun.BaseModel `bun:"table:channels,alias:ch"`

	Id          string     `bun:"id,pk"`
	UserId      string     `bun:"user_id"`
	Title       string     `bun:"title"`
	Link        string     `bun:"link"`
	Authors     string     `bun:"authors"`
	Description string     `bun:"description"`
	CreatedAt   ytime.Time `bun:"created_at"`
	UpdatedAt   ytime.Time `bun:"updated_at"`
	DeletedAt   ytime.Time `bun:"deleted_at"`
	Feed        Feed       `bun:"embed:feed_"`
}

type Feed struct {
	Url         string     `bun:"url"`
	PublishedAt ytime.Time `bun:"published_at"`
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
	CreatedAt   ytime.Time  `bun:"created_at"`
	UpdatedAt   ytime.Time  `bun:"updated_at"`
	DeletedAt   *ytime.Time `bun:"deleted_at"`
}

type File struct {
	bun.BaseModel `bun:"table:files,alias:f"`

	Id        string      `bun:"id,pk"`
	UserId    string      `bun:"user_id"`
	EpisodeId string      `bun:"episode_id"`
	UploadId  string      `bun:"upload_id"`
	UploadUrl string      `bun:"upload_url"`
	Size      int64       `bun:"size"`
	Hash      string      `bun:"hash"`
	MimeType  string      `bun:"mime_type"`
	CreatedAt ytime.Time  `bun:"created_at"`
	UpdatedAt ytime.Time  `bun:"updated_at"`
	DeletedAt *ytime.Time `bun:"deleted_at"`
}

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	Id            string `bun:"id,pk"`
	AccountNumber string `bun:"account_number"`

	CreatedAt ytime.Time  `bun:"created_at"`
	UpdatedAt ytime.Time  `bun:"updated_at"`
	DeletedAt *ytime.Time `bun:"deleted_at"`
}
