package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"mkuznets.com/go/sps/internal/auth"
)

type Api interface {
	Router() chi.Router
}

type apiImpl struct {
	authService auth.Service
	handler     Handler
}

func New(authService auth.Service, handler Handler) Api {
	return &apiImpl{authService, handler}
}

// Router ...
//
//	@title		Simple Podcast Server REST API
//	@version	0.1
//	@BasePath	/api
func (a *apiImpl) Router() chi.Router {
	r := chi.NewRouter()

	// Authenticated routes
	r.Group(func(r chi.Router) {
		r.Use(a.authService.Middleware())

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
			r.With(middleware.AllowContentType("multipart/form-data")).
				Post("/", a.handler.UploadFile)
		})

		r.Get("/users/current", nil)
	})

	// Unauthenticated routes
	r.Route("/users", func(r chi.Router) {
		r.Post("/login", a.handler.LoginUser)
		r.Post("/", a.handler.CreateUser)
	})

	return r
}
