package api

import (
	"context"
	"github.com/segmentio/ksuid"
)

type IdService interface {
	Channel(context.Context) string
	Episode(context.Context) string
	File(context.Context) string
	User(context.Context) string
}

type idService struct{}

func NewIdService() IdService {
	return &idService{}
}

func randomId() string {
	return ksuid.New().String()
}

func (s *idService) Channel(_ context.Context) string {
	return "ch_" + randomId()
}

func (s *idService) Episode(_ context.Context) string {
	return "ep_" + randomId()
}

func (s *idService) File(_ context.Context) string {
	return "file_" + randomId()
}

func (s *idService) User(_ context.Context) string {
	return "usr_" + randomId()
}
