package sps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"mkuznets.com/go/sps/docs"
	"mkuznets.com/go/sps/internal/sps/auth"
	"mkuznets.com/go/sps/internal/sps/rlog"
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
	router.Use(rlog.RequestID)
	router.Use(rlog.Logger())
	router.Use(auth.FakeAuth("usr_2K97HIyQThUf7K7AB7xdOZGAoyw"))

	swaggerSpecs := http.StripPrefix("/swagger", http.FileServer(http.FS(docs.SwaggerFiles))).ServeHTTP
	swaggerUi := httpSwagger.Handler(httpSwagger.URL("/swagger/swagger.json"))

	router.Route("/swagger", func(r chi.Router) {
		r.Get("/swagger.json", swaggerSpecs)
		r.Get("/swagger.yaml", swaggerSpecs)
		r.Handle("/*", swaggerUi)
	})
	router.Get("/swagger", http.RedirectHandler("/swagger/", http.StatusMovedPermanently).ServeHTTP)

	router.Mount("/api", s.ApiRouter)

	return http.ListenAndServe(s.Addr, router)
}
