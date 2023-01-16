package feed

import (
	"context"
	"encoding/xml"
	"mkuznets.com/go/sps/internal/rss"
	"mkuznets.com/go/sps/internal/sps/api"
	"mkuznets.com/go/sps/internal/types"
	"sort"
	"time"
)

type Controller interface {
	GetChannelIds(ctx context.Context) ([]string, error)
	Update(ctx context.Context, channelId string) error
}

type controllerImpl struct {
	store api.Store
}

func NewController(store api.Store) Controller {
	return &controllerImpl{
		store: store,
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

	//goland:noinspection HttpUrlsUsage
	podcast := &rss.Podcast{
		Version: "2.0",
		Itunes:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Channel: &rss.Channel{
			Title:         channel.Title,
			Link:          channel.Link,
			Description:   channel.Description,
			LastBuildDate: time.Now().Format(time.RFC1123Z),
			PubDate:       channel.UpdatedAt.Format(time.RFC1123Z),
			IAuthor:       channel.Authors,
		},
	}

	var items []*rss.Item
	for _, episode := range episodes {
		items = append(items, &rss.Item{
			Guid: rss.Guid{
				IsPermaLink: false,
				Text:        episode.Id,
			},
			PubDate: episode.CreatedAt.Format(time.RFC1123Z),
			Title:   episode.Title,
			Link:    episode.Link,
			Description: &rss.Description{
				Text: episode.Description,
			},
			IAuthor: episode.Authors,
			Enclosure: &rss.Enclosure{
				URL:    episode.File.Url,
				Length: episode.File.Size,
				Type:   episode.File.ContentType,
			},
		})
	}
	podcast.Channel.Items = items

	content, err := xml.Marshal(podcast)
	if err != nil {
		return err
	}

	channel.Feed = api.Feed{
		Content:     content,
		Url:         "",
		PublishedAt: types.NewTimeNow(),
	}

	if err := c.store.UpdateChannelFeeds(ctx, []*api.Channel{channel}); err != nil {
		return err
	}

	return nil
}
