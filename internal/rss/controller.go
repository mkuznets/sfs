package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"mkuznets.com/go/sfs/internal/files"
	"mkuznets.com/go/sfs/internal/store"
	"mkuznets.com/go/ytils/ytime"
	"strings"
)

type Controller interface {
	UpdateFeeds(ctx context.Context, feeds []*store.Feed) error
	BuildRss(ctx context.Context, feed *store.Feed) error
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

func (c *controllerImpl) UpdateFeeds(ctx context.Context, feeds []*store.Feed) error {
	for _, feed := range feeds {
		if err := c.BuildRss(ctx, feed); err != nil {
			return err
		}
	}
	if err := c.store.UpdateFeeds(ctx, feeds, "rss_content", "rss_content_updated_at", "rss_url", "rss_url_updated_at", "updated_at"); err != nil {
		return err
	}
	return nil
}

func (c *controllerImpl) BuildRss(ctx context.Context, feed *store.Feed) error {
	items, err := c.store.GetItems(ctx, &store.ItemFilter{FeedIds: []string{feed.Id}})
	if err != nil {
		return err
	}

	var xmlModel any
	switch feed.Type {
	case "podcast":
		xmlModel = FeedToPodcast(feed, items)
	default:
		return fmt.Errorf("%s has invalid feed type: %s", feed.Id, feed.Type)
	}

	content, err := xml.MarshalIndent(xmlModel, "", "  ")
	if err != nil {
		return err
	}

	feed.RssContent = string(content)
	feed.RssContentUpdatedAt = ytime.Now()

	path := fmt.Sprintf("rss/%s.xml", feed.Id)
	upload, err := c.fileStorage.Upload(ctx, path, strings.NewReader(feed.RssContent))
	if err != nil {
		return fmt.Errorf("failed to upload RSS feed: %w", err)
	}

	feed.RssUrl = upload.Url
	feed.RssUrlUpdatedAt = ytime.Now()

	feed.UpdatedAt = ytime.Now()

	return nil
}
