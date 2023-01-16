package api

import (
	"mkuznets.com/go/sps/internal/rss"
	"time"
)

func ChannelToPodcast(channel *Channel, episodes []*Episode) *rss.Podcast {
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

	return podcast
}
