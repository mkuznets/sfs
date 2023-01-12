package sps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"mkuznets.com/go/sps/internal/sps/auth"
	"net/http"
	"time"
)

type Server struct {
	Addr      string
	ApiRouter chi.Router
}

// Start initialises the server
func (s *Server) Start() error {
	router := chi.NewRouter()
	router.Use(middleware.Timeout(30 * time.Second))
	router.Use(middleware.Recoverer)
	router.Use(auth.FakeAuth("usr_2K97HIyQThUf7K7AB7xdOZGAoyw"))

	router.Mount("/api", s.ApiRouter)

	return http.ListenAndServe(s.Addr, router)
}
