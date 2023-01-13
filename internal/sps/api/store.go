package api

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"mkuznets.com/go/sps/internal/herror"
)

type Store interface {
	CreateChannel(ctx context.Context, channel *Channel) error
	GetChannel(ctx context.Context, id string) (*Channel, error)
	ListChannels(ctx context.Context, userId string) ([]*Channel, error)
	ListEpisodesByChannel(ctx context.Context, channelId, userId string) ([]*Episode, error)
	CreateFile(ctx context.Context, file *File) error
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

func (s *storeImpl) ListEpisodesByChannel(ctx context.Context, channelId, userId string) ([]*Episode, error) {
	return nil, nil
}

func (s *storeImpl) CreateFile(ctx context.Context, file *File) error {
	_, err := s.db.NewInsert().Model(file).Exec(ctx)
	return err
}
