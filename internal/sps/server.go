package sps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"mkuznets.com/go/sps/docs"
	"mkuznets.com/go/sps/internal/rlog"
	"mkuznets.com/go/sps/static"
	"net/http"
)

type Server struct {
	Addr      string
	ApiRouter chi.Router
	UiRouter  chi.Router
}

// Start initialises the server
func (s *Server) Start() error {
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(rlog.RequestID)
	router.Use(rlog.Logger())

	swaggerSpecs := http.StripPrefix("/swagger", http.FileServer(http.FS(docs.SwaggerFiles))).ServeHTTP
	swaggerUi := httpSwagger.Handler(httpSwagger.URL("/swagger/swagger.json"))

	router.Route("/swagger", func(r chi.Router) {
		r.Get("/swagger.json", swaggerSpecs)
		r.Get("/swagger.yaml", swaggerSpecs)
		r.Handle("/*", swaggerUi)
	})
	router.Get("/swagger", http.RedirectHandler("/swagger/", http.StatusMovedPermanently).ServeHTTP)

	router.Mount("/api", s.ApiRouter)

	router.Route("/static", func(r chi.Router) {
		r.Get("/*", http.StripPrefix("/static", http.FileServer(http.FS(static.StaticFiles))).ServeHTTP)
	})

	router.Mount("/", s.UiRouter)

	return http.ListenAndServe(s.Addr, router)
}
