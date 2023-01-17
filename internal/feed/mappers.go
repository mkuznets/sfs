package feed

import (
	"mkuznets.com/go/sps/internal/store"
	"time"
)

func ChannelToPodcast(channel *store.Channel, episodes []*store.Episode) *Podcast {
	//goland:noinspection HttpUrlsUsage
	podcast := &Podcast{
		Version: "2.0",
		Itunes:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Channel: &Channel{
			Title:         channel.Title,
			Link:          channel.Link,
			Description:   channel.Description,
			LastBuildDate: time.Now().Format(time.RFC1123Z),
			PubDate:       channel.UpdatedAt.Format(time.RFC1123Z),
			IAuthor:       channel.Authors,
		},
	}

	var items []*Item
	for _, episode := range episodes {
		items = append(items, &Item{
			Guid: Guid{
				IsPermaLink: false,
				Text:        episode.Id,
			},
			PubDate: episode.CreatedAt.Format(time.RFC1123Z),
			Title:   episode.Title,
			Link:    episode.Link,
			Description: &Description{
				Text: episode.Description,
			},
			IAuthor: episode.Authors,
			Enclosure: &Enclosure{
				URL:    episode.File.UploadUrl,
				Length: episode.File.Size,
				Type:   episode.File.MimeType,
			},
		})
	}
	podcast.Channel.Items = items

	return podcast
}
