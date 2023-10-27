package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger"
	"mkuznets.com/go/ytils/yhttp"

	"mkuznets.com/go/sfs/internal/api/swagger"
	"mkuznets.com/go/sfs/internal/auth"
)

type Api interface {
	Handler(prefix string) chi.Router
}

type apiImpl struct {
	auth    auth.Service
	handler Handler
}

func New(auth auth.Service, handler Handler) Api {
	return &apiImpl{auth, handler}
}

// Handler ...
//
//	@title						Simple Feed Service HTTP API
//	@version					0.1
//	@BasePath					/api
//	@securityDefinitions.apikey	Authentication
//	@in							header
//	@name						Authorization
func (a *apiImpl) Handler(prefix string) chi.Router {
	r := chi.NewRouter()

	r.Route(prefix, func(r chi.Router) {
		r.NotFound(func(w http.ResponseWriter, r *http.Request) {
			yhttp.Render(w, r, fmt.Errorf("HTTP 404: endpoint not found")).JSON()
		})
		r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
			yhttp.Render(w, r, fmt.Errorf("HTTP 405: %s not allowed", r.Method)).JSON()
		})

		r.Use(middleware.Recoverer)
		r.Use(AddRequestIdMiddleware)
		r.Use(AddContextLoggerMiddleware)
		r.Use(LogRequestMiddleware)

		r.Group(func(r chi.Router) {
			r.Use(a.auth.Middleware())

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
		})

		swaggerSpecs := http.FileServer(http.FS(swagger.Specs))
		r.Get("/swagger.*", http.StripPrefix(prefix, swaggerSpecs).ServeHTTP)
	})
	r.Get("/rss/{feedId}", a.handler.GetRssRedirect)

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
