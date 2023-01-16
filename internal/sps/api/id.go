package api

import "github.com/segmentio/ksuid"

type IdService interface {
	Channel() string
	Episode() string
	File() string
}

type idService struct{}

func NewIdService() IdService {
	return &idService{}
}

func (s *idService) Channel() string {
	return "ch_" + ksuid.New().String()
}

func (s *idService) Episode() string {
	return "ep_" + ksuid.New().String()
}

func (s *idService) File() string {
	return "file_" + ksuid.New().String()
}
