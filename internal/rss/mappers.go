package rss

import (
	"mkuznets.com/go/sps/internal/store"
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

	var podcastItems []*PodcastItem
	for _, episode := range items {
		podcastItems = append(podcastItems, &PodcastItem{
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
	podcast.Channel.Items = podcastItems

	return podcast
}
