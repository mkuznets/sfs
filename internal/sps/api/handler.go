package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"mkuznets.com/go/sps/internal/herror"
	"mkuznets.com/go/sps/internal/sps/auth"
	"net/http"
)

type Handler interface {
	GetChannel(w http.ResponseWriter, r *http.Request)
	CreateChannel(w http.ResponseWriter, r *http.Request)
	ListChannels(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	c Controller
}

func NewHandler(c Controller) Handler {
	return &handlerImpl{c}
}

func (h *handlerImpl) GetChannel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "channelId")

	response, err := h.c.GetChannel(r.Context(), id)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

func (h *handlerImpl) CreateChannel(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUser(r)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	var resource CreateChannelRequest
	if err := render.DecodeJSON(r.Body, &resource); err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	response, err := h.c.CreateChannel(r.Context(), user.Id(), resource)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

func (h *handlerImpl) ListChannels(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUser(r)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	response, err := h.c.ListChannels(r.Context(), user.Id())
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
