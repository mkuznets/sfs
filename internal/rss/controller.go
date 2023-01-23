package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"mkuznets.com/go/sps/internal/files"
	"mkuznets.com/go/sps/internal/store"
	"mkuznets.com/go/sps/internal/ytils/ytime"
)

type Controller interface {
	UpdateFeeds(ctx context.Context, feedIds []string) error
	BuildFeedsRss(ctx context.Context, feeds []*store.Feed) error
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

func (c *controllerImpl) UpdateFeeds(ctx context.Context, feedIds []string) error {
	feeds, err := c.store.GetFeeds(ctx, &store.FeedFilter{Ids: feedIds})
	if err != nil {
		return err
	}

	if err := c.BuildFeedsRss(ctx, feeds); err != nil {
		return err
	}

	if err := c.store.UpdateFeeds(ctx, feeds, "rss", "rss_updated_at"); err != nil {
		return err
	}

	return nil
}

func (c *controllerImpl) BuildFeedsRss(ctx context.Context, feeds []*store.Feed) error {
	for _, feed := range feeds {
		items, err := c.store.GetItems(ctx, &store.ItemFilter{FeedIds: []string{feed.Id}})
		if err != nil {
			return err
		}

		var xmlModel any
		switch feed.Type {
		case "podcast":
			xmlModel = FeedToPodcast(feed, items)
		default:
			return fmt.Errorf("invalid feed type: %s", feed.Type)
		}

		content, err := xml.Marshal(xmlModel)
		if err != nil {
			return err
		}

		feed.Rss = string(content)
		feed.RssUpdatedAt = ytime.Now()
	}

	return nil
}
