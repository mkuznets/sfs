package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"mime/multipart"
	"mkuznets.com/go/sfs/internal/user"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
	"mkuznets.com/go/sfs/internal/ytils/ynits"
	"mkuznets.com/go/sfs/internal/ytils/yrender"
	"net/http"
)

type Handler interface {
	GetFeeds(w http.ResponseWriter, r *http.Request)
	CreateFeeds(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
	CreateItems(w http.ResponseWriter, r *http.Request)
	UploadFiles(w http.ResponseWriter, r *http.Request)
	GetRss(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	c Controller
}

func NewHandler(c Controller) Handler {
	return &handlerImpl{c}
}

func (h *handlerImpl) GetFeeds(w http.ResponseWriter, r *http.Request) {
	usr := yerr.Must(user.Get(r))

	req, err := yrender.DecodeJson[GetFeedsRequest](r.Body)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	response, err := h.c.GetFeeds(r.Context(), &req, usr)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	yrender.New(w, r, response).JSON()
}

func (h *handlerImpl) CreateFeeds(w http.ResponseWriter, r *http.Request) {
	usr := yerr.Must(user.Get(r))

	req, err := yrender.DecodeJson[CreateFeedsRequest](r.Body)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	response, err := h.c.CreateFeeds(r.Context(), &req, usr)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	yrender.New(w, r, response).JSON()
}

func (h *handlerImpl) GetItems(w http.ResponseWriter, r *http.Request) {
	usr := yerr.Must(user.Get(r))

	req, err := yrender.DecodeJson[GetItemsRequest](r.Body)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	response, err := h.c.GetItems(r.Context(), &req, usr)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	yrender.New(w, r, response).JSON()
}

func (h *handlerImpl) CreateItems(w http.ResponseWriter, r *http.Request) {
	usr := yerr.Must(user.Get(r))

	req, err := yrender.DecodeJson[CreateItemsRequest](r.Body)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	response, err := h.c.CreateItems(r.Context(), &req, usr)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	yrender.New(w, r, response).JSON()
}

func (h *handlerImpl) UploadFiles(w http.ResponseWriter, r *http.Request) {
	usr := yerr.Must(user.Get(r))
	ctx := r.Context()

	if err := r.ParseMultipartForm(512 * ynits.MiB); err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	files := make([]multipart.File, 0)
	for _, fileHeader := range r.MultipartForm.File["file"] {
		file, err := fileHeader.Open()
		if err != nil {
			yrender.New(w, r, err).JSON()
			return
		}
		files = append(files, file)
	}

	if len(files) == 0 {
		yrender.New(w, r, yerr.Invalid("no files provided")).JSON()
		return
	}

	defer func(fs []multipart.File) {
		for _, f := range fs {
			if err := f.Close(); err != nil {
				log.Warn().Err(err).Msg("failed to close file")
			}
		}
	}(files)

	response, err := h.c.UploadFiles(ctx, files, usr)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	yrender.New(w, r, response).JSON()
}

func (h *handlerImpl) GetRss(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "feedId")
	usr := yerr.Must(user.Get(r))

	response, err := h.c.GetRss(r.Context(), id, usr)
	if err != nil {
		yrender.New(w, r, err).JSON()
		return
	}

	yrender.New(w, r, response).XML()
}
