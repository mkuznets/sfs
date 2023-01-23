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
//	@title						Simple Feed Service HTTP API
//	@version					0.1
//	@BasePath					/api
//	@securityDefinitions.apikey	Authentication
//	@in							header
//	@name						Authorization
func (a *apiImpl) Router() chi.Router {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Route("/feeds", func(r chi.Router) {
			r.Post("/get", a.handler.GetFeeds)
			r.Post("/create", a.handler.CreateFeeds)
		})
		r.Route(`/items`, func(r chi.Router) {
			r.Post(`/get`, a.handler.GetItems)
			r.Post(`/create`, a.handler.CreateItems)
		})
		r.Route("/files", func(r chi.Router) {
			r.With(middleware.AllowContentType("multipart/form-data")).
				Post("/upload", a.handler.UploadFiles)
		})
		r.Route("/rss", func(r chi.Router) {
			r.Get("/{feedId}", a.handler.GetRss)
		})
	})

	return r
}
