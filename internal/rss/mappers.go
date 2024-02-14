package rss

import (
	"time"

	"mkuznets.com/go/sfs/internal/feedstore"
)

func FeedToPodcast(feed *feedstore.Feed, items []*feedstore.Item) *Podcast {
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

	podcastItems := make([]*PodcastItem, 0, len(items))

	for _, item := range items {
		pubDate := item.PublishedAt
		if pubDate.IsZero() {
			pubDate = item.UpdatedAt
		}

		podcastItems = append(podcastItems, &PodcastItem{
			Guid: Guid{
				IsPermaLink: false,
				Text:        item.Id,
			},
			PubDate: pubDate.Format(time.RFC1123Z),
			Title:   item.Title,
			Link:    item.Link,
			Description: &Description{
				Text: item.Description,
			},
			IAuthor: item.Authors,
			Enclosure: &Enclosure{
				URL:    item.File.UploadUrl,
				Length: item.File.Size,
				Type:   item.File.MimeType,
			},
		})
	}

	podcast.Channel.Items = podcastItems

	return podcast
}
