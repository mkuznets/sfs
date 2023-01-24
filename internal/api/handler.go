package api

import (
	"github.com/go-chi/chi"
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
	u := user.MustGet(r)

	req, err := yrender.DecodeJson[GetFeedsRequest](r.Body)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	response, err := h.c.GetFeeds(r.Context(), &req, u)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) CreateFeeds(w http.ResponseWriter, r *http.Request) {
	u := user.MustGet(r)

	req, err := yrender.DecodeJson[CreateFeedsRequest](r.Body)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	response, err := h.c.CreateFeeds(r.Context(), &req, u)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) GetItems(w http.ResponseWriter, r *http.Request) {
	u := user.MustGet(r)

	req, err := yrender.DecodeJson[GetItemsRequest](r.Body)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	response, err := h.c.GetItems(r.Context(), &req, u)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) CreateItems(w http.ResponseWriter, r *http.Request) {
	u := user.MustGet(r)

	req, err := yrender.DecodeJson[CreateItemsRequest](r.Body)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	response, err := h.c.CreateItems(r.Context(), &req, u)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) UploadFiles(w http.ResponseWriter, r *http.Request) {
	u := user.MustGet(r)

	ctx := r.Context()

	if err := r.ParseMultipartForm(512 * ynits.MiB); err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	files := make([]multipart.File, 0)
	for _, fileHeader := range r.MultipartForm.File["file"] {
		file, err := fileHeader.Open()
		if err != nil {
			yrender.JsonErr(w, r, err)
			return
		}
		files = append(files, file)
	}

	if len(files) == 0 {
		yrender.JsonErr(w, r, yerr.Invalid("no files provided"))
		return
	}

	defer func(fs []multipart.File) {
		for _, f := range fs {
			if err := f.Close(); err != nil {
				log.Warn().Err(err).Msg("failed to close file")
			}
		}
	}(files)

	response, err := h.c.UploadFiles(ctx, files, u)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	yrender.Json(w, r, http.StatusOK, response)
}

func (h *handlerImpl) GetRss(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "feedId")

	response, err := h.c.GetRss(r.Context(), id)
	if err != nil {
		yrender.JsonErr(w, r, err)
		return
	}

	yrender.Rss(w, r, http.StatusOK, response)
}
