package api

import (
	"github.com/go-chi/chi"
)

// NewRouter ...
//
//	@title		Simple Podcast Server REST API
//	@version	0.1
//	@BasePath	/api
func NewRouter(handler Handler) chi.Router {
	r := chi.NewRouter()
	r.Route("/channels", func(r chi.Router) {
		r.Route("/{channelId}", func(r chi.Router) {
			r.Get(`/`, handler.GetChannel)
			r.Route(`/episodes`, func(r chi.Router) {
				r.Post(`/`, handler.CreateEpisode)
				r.Get(`/`, handler.ListEpisodes)
			})
		})
		r.Get("/", handler.ListChannels)
		r.Post("/", handler.CreateChannel)
	})
	r.Route("/files", func(r chi.Router) {
		r.Post("/", handler.UploadFile)
	})

	return r
}
