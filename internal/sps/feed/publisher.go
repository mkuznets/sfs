package feed

import "net/url"

type Publisher interface {
	Publish(feedContent []byte) (*url.URL, error)
}
