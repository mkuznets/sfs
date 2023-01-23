package ui

import (
	"github.com/go-chi/chi"
	"mkuznets.com/go/sps/internal/auth"
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"net/http"
)

type Ui interface {
	Router() chi.Router
}

type uiImpl struct {
	authService auth.Service
	handler     Handler
}

func New(authService auth.Service, handler Handler) Ui {
	return &uiImpl{authService, handler}
}

func (a *uiImpl) Router() chi.Router {
	r := chi.NewRouter()
	r.Use(a.authService.Middleware())
	r.Use(ensureUser)

	r.Route("/", func(r chi.Router) {
		r.Post("/login", a.handler.Login)
		r.Get("/logout", a.handler.Logout)
		r.Post("/logout", a.handler.Logout)
		r.Get("/", a.handler.Index)
		r.Route("/podcasts", func(r chi.Router) {
			r.Get("/", http.RedirectHandler("/", http.StatusFound).ServeHTTP)
			r.Get("/{id}", a.handler.Episodes)
			r.Post("/delete", a.handler.PodcastsDelete)
			//	r.Get("/info", a.handler.PodcastInfo)
			//	r.Get("/breadcrumbs", a.handler.PodcastBreadcrumbs)
		})
	})

	return r
}

// middleware
func ensureUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isHx := r.Header.Get("HX-Request") == "true"

		user, err := auth.GetUser(r)
		if err != nil {
			yerr.RenderJson(w, r, err)
			return
		}

		if user == nil && r.URL.Path != "/login" {
			if isHx {
				w.Header().Set("HX-Redirect", "/")
				w.WriteHeader(200)
				return
			} else if r.URL.Path != "/" {
				http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			}
		}
		next.ServeHTTP(w, r)
	})
}
