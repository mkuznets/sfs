package rss

import (
	"encoding/xml"
)

type Podcast struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Itunes  string   `xml:"xmlns:itunes,attr"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title         string         `xml:"title"`
	Link          string         `xml:"link,omitempty"`
	Description   string         `xml:"description,omitempty"`
	LastBuildDate string         `xml:"lastBuildDate"`
	PubDate       string         `xml:"pubDate"`
	IAuthor       string         `xml:"itunes:author,omitempty"`
	Items         []*PodcastItem `xml:"item"`
}

type PodcastItem struct {
	Guid        Guid         `xml:"guid"`
	Title       string       `xml:"title"`
	Link        string       `xml:"link"`
	Description *Description `xml:"description"`
	PubDate     string       `xml:"pubDate"`
	Enclosure   *Enclosure   `xml:"enclosure"`
	IDuration   *string      `xml:"itunes:duration"`
	IAuthor     string       `xml:"itunes:author"`
}

type Guid struct {
	IsPermaLink bool   `xml:"isPermaLink,attr"`
	Text        string `xml:",chardata"`
}

type Description struct {
	Text string `xml:",cdata"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}
