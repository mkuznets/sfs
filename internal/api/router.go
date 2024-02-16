package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"

	"mkuznets.com/go/sfs/internal/api/swagger"
	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/render"
)

// NewRouter ...
//
//	@title						Simple Feed Service HTTP API
//	@version					0.1
//	@BasePath					/api
//	@securityDefinitions.apikey	Authentication
//	@in							header
//	@name						Authorization
func NewRouter(prefix string, authService auth.Service, apiService *Service) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Use(AddContextLoggerMiddleware)
	r.Use(LogRequestMiddleware)

	r.Use(middleware.RealIP)
	r.Use(AddRequestIDMiddleware)

	r.Route(prefix, func(r chi.Router) {
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			render.New(w, r, fmt.Errorf("HTTP 404: endpoint not found")).JSON()
		})
		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			render.New(w, r, fmt.Errorf("HTTP 405: %s not allowed", r.Method)).JSON()
		})

		r.Group(func(r chi.Router) {
			r.Use(authService.Middleware())
			r.Use(LogUserMiddleware)

			r.Route("/feeds", func(r chi.Router) {
				r.Post("/get", apiService.GetFeeds)
				r.Post("/create", apiService.CreateFeeds)
			})
			r.Route(`/items`, func(r chi.Router) {
				r.Post(`/get`, apiService.GetItems)
				r.Post(`/create`, apiService.CreateItems)
			})
			r.Route("/files", func(r chi.Router) {
				r.With(middleware.AllowContentType("multipart/form-data")).
					Post("/upload", apiService.UploadFiles)
			})
		})

		swaggerSpecs := http.FileServer(http.FS(swagger.Specs))
		r.Get("/swagger.*", http.StripPrefix(prefix, swaggerSpecs).ServeHTTP)
	})
	r.Get("/rss/{feedId}", apiService.GetRssRedirect)

	swaggerUi := httpSwagger.Handler(
		httpSwagger.URL(prefix+"/swagger.json"),
		httpSwagger.UIConfig(map[string]string{
			"displayOperationId":       "true",
			"deepLinking":              "true",
			"defaultModelsExpandDepth": "-1",
			"defaultModelExpandDepth":  "5",
			"displayRequestDuration":   "true",
		}))
	r.Get("/swagger/*", swaggerUi)
	r.Get("/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently).ServeHTTP)

	return r
}
