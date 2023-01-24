package sps

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"mkuznets.com/go/sps/internal/api/swagger"
	"mkuznets.com/go/sps/internal/ytils/ylog"
	"mkuznets.com/go/sps/internal/ytils/yreq"
	"net/http"
)

type Server struct {
	Addr      string
	ApiRouter chi.Router
}

// Start initialises the server
func (s *Server) Start() error {
	router := chi.NewRouter()
	router.Use(middleware.StripSlashes)
	router.Use(middleware.Recoverer)
	router.Use(yreq.RequestId)
	router.Use(ylog.ContextLogger)

	router.Mount("/api", s.ApiRouter)

	specs := http.FileServer(http.FS(swagger.Specs))
	router.Get("/api/swagger.*", http.StripPrefix("/api", specs).ServeHTTP)

	swaggerUi := httpSwagger.Handler(
		httpSwagger.URL("/api/swagger.json"),
		httpSwagger.UIConfig(map[string]string{
			"displayOperationId":       "true",
			"deepLinking":              "true",
			"defaultModelsExpandDepth": "-1",
			"defaultModelExpandDepth":  "5",
			"displayRequestDuration":   "true",
		}))

	router.Mount("/swagger", swaggerUi)

	return http.ListenAndServe(s.Addr, router)
}
