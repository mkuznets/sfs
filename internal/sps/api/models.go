package api

import (
	"github.com/segmentio/ksuid"
	"github.com/uptrace/bun"
	"mkuznets.com/go/sps/internal/types"
)

type Channel struct {
	bun.BaseModel `bun:"table:channels,alias:ch"`

	Id          string      `bun:"id"`
	UserId      string      `bun:"user_id"`
	Title       string      `bun:"title"`
	Link        string      `bun:"link"`
	Authors     string      `bun:"authors"`
	Description string      `bun:"description"`
	CreatedAt   types.Time  `bun:"created_at"`
	UpdatedAt   types.Time  `bun:"updated_at"`
	DeletedAt   *types.Time `bun:"deleted_at"`
}

func RandomChannelId() string {
	return "ch_" + ksuid.New().String()
}

type Episode struct {
	Id        string      `bun:"id"`
	ChannelId string      `bun:"channel_id"`
	Title     string      `bun:"title"`
	CreatedAt types.Time  `bun:"created_at"`
	UpdatedAt types.Time  `bun:"updated_at"`
	DeletedAt *types.Time `bun:"deleted_at"`
}

func RandomEpisodeId() string {
	return "ep_" + ksuid.New().String()
}

type File struct {
	bun.BaseModel `bun:"table:files,alias:f"`

	Id          string      `bun:"id"`
	UserId      string      `bun:"user_id"`
	Url         string      `bun:"url"`
	Size        int64       `bun:"size"`
	ContentType string      `bun:"content_type"`
	CreatedAt   types.Time  `bun:"created_at"`
	UpdatedAt   types.Time  `bun:"updated_at"`
	DeletedAt   *types.Time `bun:"deleted_at"`
}

func RandomFileId() string {
	return "file_" + ksuid.New().String()
}
