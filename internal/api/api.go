package api

import "github.com/go-chi/chi"

type Api interface {
	Router() chi.Router
}

type apiImpl struct {
	handler Handler
}

func New(handler Handler) Api {
	return &apiImpl{handler}
}

// Router ...
//
//	@title		Simple Podcast Server REST API
//	@version	0.1
//	@BasePath	/api
func (a *apiImpl) Router() chi.Router {
	r := chi.NewRouter()
	r.Route("/channels", func(r chi.Router) {
		r.Route("/{channelId}", func(r chi.Router) {
			r.Get(`/`, a.handler.GetChannel)
			r.Route(`/episodes`, func(r chi.Router) {
				r.Post(`/`, a.handler.CreateEpisode)
				r.Get(`/`, a.handler.ListEpisodes)
			})
		})
		r.Get("/", a.handler.ListChannels)
		r.Post("/", a.handler.CreateChannel)
	})
	r.Route("/files", func(r chi.Router) {
		r.Post("/", a.handler.UploadFile)
	})

	return r
}
