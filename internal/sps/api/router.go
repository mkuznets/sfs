package api

import (
	"github.com/go-chi/chi"
)

func NewRouter(handler Handler) chi.Router {
	r := chi.NewRouter()
	r.Route("/channels", func(r chi.Router) {
		r.Get(`/{channelId:\w+}`, handler.GetChannel)
		r.Get("/", handler.ListChannels)
		r.Post("/", handler.CreateChannel)
	})

	return r
}
