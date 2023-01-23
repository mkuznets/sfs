package ui

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"mkuznets.com/go/sps/internal/api"
	"mkuznets.com/go/sps/internal/auth"
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"net/http"
	"time"
)

type Handler interface {
	Index(w http.ResponseWriter, r *http.Request)
	PodcastsDelete(w http.ResponseWriter, r *http.Request)
	Episodes(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	c  Controller
	ac api.Controller
}

func NewHandler(c Controller, ac api.Controller) Handler {
	return &handlerImpl{
		c:  c,
		ac: ac,
	}
}

func isHx(r *http.Request) bool {
	return r.Header.Get("HX-Request") == "true"
}

func (h *handlerImpl) Login(w http.ResponseWriter, r *http.Request) {
	login := r.FormValue("login")

	token, err := h.ac.Login(r.Context(), &api.LoginRequest{AccountNumber: login})
	if err != nil {
		w.WriteHeader(200)
		w.Write([]byte(err.Error()))
		return
	}

	// Set session cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "JWT",
		Value:   token,
		Expires: time.Now().Add(30 * 24 * time.Hour),
		Path:    "/",
	})

	if isHx(r) {
		w.Header().Set("HX-Refresh", "true")
		w.WriteHeader(200)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *handlerImpl) Logout(w http.ResponseWriter, r *http.Request) {
	// Delete JWT cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "JWT",
		Value:   "",
		Expires: time.Now(),
		MaxAge:  -1,
		Path:    "/",
	})

	if isHx(r) {
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(200)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *handlerImpl) render(w http.ResponseWriter, r *http.Request, content string) {
	user, err := auth.GetUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	if isHx(r) {
		render.HTML(w, r, content)
		return
	}

	page := Page{Content: content}
	if user != nil {
		page.User = user.AccountNumber()
	}

	output, err := h.c.Page(r.Context(), page)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	render.HTML(w, r, output)
}

func (h *handlerImpl) Index(w http.ResponseWriter, r *http.Request) {
	var (
		output string
		err    error
	)

	user, err := auth.GetUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	if user == nil {
		if isHx(r) {
			w.Header().Set("HX-Redirect", "/")
			w.WriteHeader(200)
			return
		}
		output, err = h.c.GetLoginContent(r.Context())

	} else {
		output, err = h.c.GetPodcastsContent(r.Context(), user.Id())
	}

	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	h.render(w, r, output)
}

func (h *handlerImpl) Episodes(w http.ResponseWriter, r *http.Request) {
	user, err := auth.RequireUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}
	channelId := chi.URLParam(r, "id")

	output, err := h.c.GetEpisodesContent(r.Context(), user.Id(), channelId)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	h.render(w, r, output)
}

func (h *handlerImpl) PodcastsDelete(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	log.Ctx(r.Context()).Debug().Strs("values", r.PostForm["id"]).Msg("PodcastsDelete")

	if isHx(r) {
		w.Header().Set("HX-Trigger", "tableRefresh")
		w.WriteHeader(200)
		return
	}
}
