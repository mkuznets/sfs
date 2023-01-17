package feed

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"mkuznets.com/go/sps/internal/files"
	"mkuznets.com/go/sps/internal/store"
	"mkuznets.com/go/sps/internal/ytils/ytime"
	"sort"
)

type Controller interface {
	GetChannelIds(ctx context.Context) ([]string, error)
	Update(ctx context.Context, channelId string) error
}

type controllerImpl struct {
	store       store.Store
	fileStorage files.Storage
}

func NewController(store store.Store, fileStorage files.Storage) Controller {
	return &controllerImpl{
		store:       store,
		fileStorage: fileStorage,
	}
}

func (c *controllerImpl) GetChannelIds(ctx context.Context) ([]string, error) {
	return c.store.GetChannelsIdsToUpdateFeeds(ctx)
}

func (c *controllerImpl) Update(ctx context.Context, channelId string) error {
	channel, err := c.store.GetChannel(ctx, channelId)
	if err != nil {
		return err
	}

	episodes, err := c.store.ListEpisodesWithFiles(ctx, channelId)
	if err != nil {
		return err
	}
	sort.Slice(episodes, func(i, j int) bool {
		return episodes[i].CreatedAt.After(episodes[j].CreatedAt.Time)
	})

	podcast := ChannelToPodcast(channel, episodes)

	content, err := xml.Marshal(podcast)
	if err != nil {
		return err
	}

	path := fmt.Sprintf("rss/%s/feed.xml", channel.Id)
	upload, err := c.fileStorage.Upload(ctx, path, bytes.NewReader(content))
	if err != nil {
		return err
	}

	channel.Feed = store.Feed{
		Url:         upload.Url,
		PublishedAt: ytime.Now(),
	}

	if err := c.store.UpdateChannelFeeds(ctx, []*store.Channel{channel}); err != nil {
		return err
	}

	return nil
}
