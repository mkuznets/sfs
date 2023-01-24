package rss

import (
	"mkuznets.com/go/sfs/internal/store"
	"mkuznets.com/go/sfs/internal/ytils/yslice"
	"time"
)

func FeedToPodcast(feed *store.Feed, items []*store.Item) *Podcast {
	//goland:noinspection HttpUrlsUsage
	podcast := &Podcast{
		Version: "2.0",
		Itunes:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
		Channel: &Channel{
			Title:         feed.Title,
			Link:          feed.Link,
			Description:   feed.Description,
			LastBuildDate: time.Now().Format(time.RFC1123Z),
			PubDate:       feed.UpdatedAt.Format(time.RFC1123Z),
			IAuthor:       feed.Authors,
		},
	}

	podcast.Channel.Items = yslice.Map(items, func(i *store.Item) *PodcastItem {
		return &PodcastItem{
			Guid: Guid{
				IsPermaLink: false,
				Text:        i.Id,
			},
			PubDate: i.CreatedAt.Format(time.RFC1123Z),
			Title:   i.Title,
			Link:    i.Link,
			Description: &Description{
				Text: i.Description,
			},
			IAuthor: i.Authors,
			Enclosure: &Enclosure{
				URL:    i.File.UploadUrl,
				Length: i.File.Size,
				Type:   i.File.MimeType,
			},
		}
	})

	return podcast
}
