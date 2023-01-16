package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"mime/multipart"
	"mkuznets.com/go/sps/internal/herror"
	"mkuznets.com/go/sps/internal/sps/auth"
	"mkuznets.com/go/sps/internal/sps/rlog"
	"net/http"
)

const (
	MiB = 1024 * 1024
)

type Handler interface {
	GetChannel(w http.ResponseWriter, r *http.Request)
	GetFeed(w http.ResponseWriter, r *http.Request)
	CreateChannel(w http.ResponseWriter, r *http.Request)
	ListChannels(w http.ResponseWriter, r *http.Request)
	UploadFile(w http.ResponseWriter, r *http.Request)
	CreateEpisode(w http.ResponseWriter, r *http.Request)
	ListEpisodes(w http.ResponseWriter, r *http.Request)
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

func (h *handlerImpl) GetFeed(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "channelId")

	content, err := h.c.GetChannelFeed(r.Context(), id)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	w.Header().Add("Content-Type", "application/rss+xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(content); err != nil {
		log.Debug().Err(err).Msg("failed to write response")
	}
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

func (h *handlerImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUser(r)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	ctx := r.Context()

	if err := r.ParseMultipartForm(512 * MiB); err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err == http.ErrMissingFile {
		herror.RenderJson(w, r, herror.Validation("no file provided"))
		return
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(file)

	rlog.L(ctx).Debug().
		Int64("size", header.Size).
		Str("name", header.Filename).
		Msg("uploading file")

	response, err := h.c.UploadFile(ctx, user.Id(), file)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

func (h *handlerImpl) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUser(r)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	channelId := chi.URLParam(r, "channelId")

	var resource CreateEpisodeRequest
	if err := render.DecodeJSON(r.Body, &resource); err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	response, err := h.c.CreateEpisode(r.Context(), user.Id(), channelId, &resource)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}

func (h *handlerImpl) ListEpisodes(w http.ResponseWriter, r *http.Request) {
	user, err := auth.GetUser(r)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}
	channelId := chi.URLParam(r, "channelId")

	response, err := h.c.ListEpisodes(r.Context(), user.Id(), channelId)
	if err != nil {
		herror.RenderJson(w, r, err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response)
}
