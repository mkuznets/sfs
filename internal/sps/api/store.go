package api

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"mkuznets.com/go/sps/internal/herror"
	"mkuznets.com/go/sps/internal/types"
)

type Store interface {
	CreateChannel(ctx context.Context, channel *Channel) error
	GetChannel(ctx context.Context, id string) (*Channel, error)
	ListChannels(ctx context.Context, userId string) ([]*Channel, error)
	ListEpisodesWithFiles(ctx context.Context, channelId string) ([]*Episode, error)
	CreateEpisode(ctx context.Context, episode *Episode) error
	CreateFile(ctx context.Context, file *File) error
	GetFile(ctx context.Context, id string) (*File, error)
	GetChannelsIdsToUpdateFeeds(ctx context.Context) ([]string, error)
	UpdateChannelFeeds(ctx context.Context, channels []*Channel) error
}

type storeImpl struct {
	db *bun.DB
}

func NewStore(db *bun.DB) Store {
	return &storeImpl{db: db}
}

func (s *storeImpl) CreateChannel(ctx context.Context, channel *Channel) error {
	_, err := s.db.NewInsert().Model(channel).Exec(ctx)
	return err
}

func (s *storeImpl) GetChannel(ctx context.Context, id string) (*Channel, error) {
	var channel Channel
	err := s.db.NewSelect().Model(&channel).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, herror.NotFound("channel not found")
		}
		return nil, err
	}
	return &channel, nil
}

func (s *storeImpl) ListChannels(ctx context.Context, userId string) ([]*Channel, error) {
	channels := make([]*Channel, 0)
	err := s.db.NewSelect().Model(&channels).
		Where("user_id = ?", userId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return channels, nil
}

func (s *storeImpl) ListEpisodesWithFiles(ctx context.Context, channelId string) ([]*Episode, error) {
	episodes := make([]*Episode, 0)
	err := s.db.NewSelect().Model(&episodes).
		Relation("File").
		Where("channel_id = ?", channelId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return episodes, nil
}

func (s *storeImpl) CreateEpisode(ctx context.Context, episode *Episode) error {
	_, err := s.db.NewInsert().Model(episode).Exec(ctx)
	if err != nil {
		return err
	}

	_, err = s.db.NewUpdate().
		Model(&Channel{}).
		Set("updated_at = ?", types.NewTimeNow()).
		Where("id = ?", episode.ChannelId).
		Exec(ctx)
	if err != nil {
		return err
	}

	return err
}

func (s *storeImpl) CreateFile(ctx context.Context, file *File) error {
	_, err := s.db.NewInsert().Model(file).Exec(ctx)
	return err
}

func (s *storeImpl) GetFile(ctx context.Context, id string) (*File, error) {
	var file File
	err := s.db.NewSelect().Model(&file).
		Where("id = ?", id).
		Scan(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, herror.NotFound("file not found")
		}
		return nil, err
	}
	return &file, nil
}

func (s *storeImpl) GetChannelsIdsToUpdateFeeds(ctx context.Context) ([]string, error) {
	ids := make([]string, 0)
	err := s.db.NewSelect().
		ColumnExpr("id").
		Model((*Channel)(nil)).
		Where("feed_published_at IS NULL OR feed_published_at < updated_at").
		Scan(ctx, &ids)
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (s *storeImpl) UpdateChannelFeeds(ctx context.Context, channels []*Channel) error {
	values := s.db.NewValues(&channels).Column("id", "feed_content", "feed_published_at", "feed_url")
	_, err := s.db.NewUpdate().
		With("_data", values).
		Model((*Channel)(nil)).
		TableExpr("_data").
		Set("feed_content = _data.feed_content").
		Set("feed_published_at = _data.feed_published_at").
		Set("feed_url = _data.feed_url").
		Where("ch.id = _data.id").
		Exec(ctx)
	return err
}
