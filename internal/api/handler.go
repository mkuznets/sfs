package api

import (
	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
	"mime/multipart"
	"mkuznets.com/go/sps/internal/auth"
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"mkuznets.com/go/sps/internal/ytils/ynits"
	"mkuznets.com/go/sps/internal/ytils/yrender"
	"net/http"
)

type Handler interface {
	GetChannel(w http.ResponseWriter, r *http.Request)
	CreateChannel(w http.ResponseWriter, r *http.Request)
	ListChannels(w http.ResponseWriter, r *http.Request)
	UploadFile(w http.ResponseWriter, r *http.Request)
	CreateEpisode(w http.ResponseWriter, r *http.Request)
	ListEpisodes(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
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
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) CreateChannel(w http.ResponseWriter, r *http.Request) {
	user, err := auth.RequireUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	var resource CreateChannelRequest
	if err := yrender.DecodeJson(r.Body, &resource); err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	response, err := h.c.CreateChannel(r.Context(), user.Id(), resource)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) ListChannels(w http.ResponseWriter, r *http.Request) {
	user, err := auth.RequireUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	response, err := h.c.ListChannels(r.Context(), user.Id())
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) UploadFile(w http.ResponseWriter, r *http.Request) {
	user, err := auth.RequireUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	ctx := r.Context()

	if err := r.ParseMultipartForm(512 * ynits.MiB); err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	file, header, err := r.FormFile("file")
	if err == http.ErrMissingFile {
		yerr.RenderJson(w, r, yerr.Validation("no file provided"))
		return
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(file)

	log.Ctx(ctx).Debug().
		Int64("size", header.Size).
		Str("name", header.Filename).
		Msg("uploading file")

	response, err := h.c.UploadFile(ctx, user.Id(), file)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) CreateEpisode(w http.ResponseWriter, r *http.Request) {
	user, err := auth.RequireUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	channelId := chi.URLParam(r, "channelId")

	var resource CreateEpisodeRequest
	if err := yrender.DecodeJson(r.Body, &resource); err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	response, err := h.c.CreateEpisode(r.Context(), user.Id(), channelId, &resource)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) ListEpisodes(w http.ResponseWriter, r *http.Request) {
	user, err := auth.RequireUser(r)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}
	channelId := chi.URLParam(r, "channelId")

	response, err := h.c.ListEpisodes(r.Context(), user.Id(), channelId)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	response, err := h.c.CreateUser(r.Context())
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) LoginUser(w http.ResponseWriter, r *http.Request) {
	var resource LoginRequest
	if err := yrender.DecodeJson(r.Body, &resource); err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	token, err := h.c.Login(r.Context(), &resource)
	if err != nil {
		yerr.RenderJson(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, token)
}
