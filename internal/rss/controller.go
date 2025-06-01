package rss

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"time"

	"mkuznets.com/go/sfs/internal/feedstore"
	"mkuznets.com/go/sfs/internal/filestore"
	"mkuznets.com/go/sfs/internal/mtime"
)

type Controller interface {
	UpdateFeeds(ctx context.Context, feeds []*feedstore.Feed) error
	BuildRss(ctx context.Context, feed *feedstore.Feed) error
	GetFeedContent(ctx context.Context, feed *feedstore.Feed) (*FeedContent, error)
}
type controllerImpl struct {
	feedStore feedstore.FeedStore
	fileStore filestore.FileStore
}

func NewController(store feedstore.FeedStore, fileStorage filestore.FileStore) Controller {
	return &controllerImpl{
		feedStore: store,
		fileStore: fileStorage,
	}
}

func (c *controllerImpl) UpdateFeeds(ctx context.Context, feeds []*feedstore.Feed) error {
	for _, feed := range feeds {
		if err := c.BuildRss(ctx, feed); err != nil {
			return err
		}
	}
	if err := c.feedStore.UpdateFeeds(ctx, feeds, "rss_content", "rss_content_updated_at", "rss_url", "rss_url_updated_at", "updated_at"); err != nil {
		return err
	}
	return nil
}

func (c *controllerImpl) BuildRss(ctx context.Context, feed *feedstore.Feed) error {
	items, err := c.feedStore.GetItems(ctx, &feedstore.ItemFilter{FeedIds: []string{feed.Id}})
	if err != nil {
		return err
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].PublishedAt.After(items[j].PublishedAt.Time)
	})

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

	feed.RssContentUpdatedAt = mtime.Now()

	path := fmt.Sprintf("rss/%s.xml", feed.Id)
	upload, err := c.fileStore.Upload(ctx, path, bytes.NewReader(content))
	if err != nil {
		return fmt.Errorf("failed to upload RSS feed: %w", err)
	}

	feed.RssUrl = upload.URL
	feed.RssUrlUpdatedAt = mtime.Now()

	feed.UpdatedAt = mtime.Now()

	return nil
}

type FeedContent struct {
	LastModified time.Time
	ETag         string
	Reader       io.ReadSeeker
}

func (c *controllerImpl) GetFeedContent(ctx context.Context, feed *feedstore.Feed) (*FeedContent, error) {
	items, err := c.feedStore.GetItems(ctx, &feedstore.ItemFilter{FeedIds: []string{feed.Id}})
	if err != nil {
		return nil, err
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].PublishedAt.After(items[j].PublishedAt.Time)
	})

	var xmlModel any
	switch feed.Type {
	case "podcast":
		xmlModel = FeedToPodcast(feed, items)
	default:
		return nil, fmt.Errorf("%s has invalid feed type: %s", feed.Id, feed.Type)
	}

	content, err := xml.MarshalIndent(xmlModel, "", "  ")
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write(content)
	etag := fmt.Sprintf("%x", h.Sum(nil))

	result := &FeedContent{
		LastModified: feed.UpdatedAt.Time,
		ETag:         etag,
		Reader:       bytes.NewReader(content),
	}

	return result, nil
}
