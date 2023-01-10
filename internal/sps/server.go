package sps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

type Server struct {
	Addr string
	Api  *Api
}

// Start initialises the server
func (s *Server) Start() error {
	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(middleware.Recoverer)

	router.Route("/channels", func(r chi.Router) {
		r.Get("/", s.Api.ListChannels)
	})

	return http.ListenAndServe(s.Addr, router)
}
