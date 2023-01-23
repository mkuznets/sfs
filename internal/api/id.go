package api

import (
	"context"
	"github.com/segmentio/ksuid"
)

type IdService interface {
	Feed(context.Context) string
	Item(context.Context) string
	File(context.Context) string
	Rss(context.Context) string
}

type idService struct{}

func NewIdService() IdService {
	return &idService{}
}

func randomId() string {
	return ksuid.New().String()
}

func (s *idService) Feed(_ context.Context) string {
	return "feed_" + randomId()
}

func (s *idService) Item(_ context.Context) string {
	return "item_" + randomId()
}

func (s *idService) File(_ context.Context) string {
	return "file_" + randomId()
}

func (s *idService) Rss(_ context.Context) string {
	return "rss_" + randomId()
}
