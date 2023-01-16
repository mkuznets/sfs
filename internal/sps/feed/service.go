package feed

import (
	"context"
	"github.com/rs/zerolog/log"
	"time"
)

type Service interface {
	Start(ctx context.Context)
}

type serviceImpl struct {
	c Controller
}

func NewService(controller Controller) Service {
	return &serviceImpl{
		c: controller,
	}
}

func (s *serviceImpl) Start(ctx context.Context) {
	for {
		func() {
			ids, err := s.c.GetChannelIds(ctx)
			if err != nil {
				log.Err(err).Msg("failed to get channel ids")
			}
			log.Debug().Strs("ids", ids).Msg("updating channel feeds")

			for _, id := range ids {
				if err := s.c.Update(ctx, id); err != nil {
					log.Err(err).Msg("failed to update feed")
				}
			}
		}()
		time.Sleep(5 * time.Second)
	}
}
