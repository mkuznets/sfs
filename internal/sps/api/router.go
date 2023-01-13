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
		r.Get(`/{channelId:\w+}`, handler.GetChannel)
		r.Get("/", handler.ListChannels)
		r.Post("/", handler.CreateChannel)
	})

	r.Route("/files", func(r chi.Router) {
		r.Post("/", handler.UploadFile)
	})

	return r
}
