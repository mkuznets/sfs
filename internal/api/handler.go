package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mkuznets.com/go/ytils/y"
	"mkuznets.com/go/ytils/yhttp"

	"mkuznets.com/go/sfs/internal/slogger"
	"mkuznets.com/go/sfs/internal/user"
)

type Handler interface {
	GetFeeds(w http.ResponseWriter, r *http.Request)
	CreateFeeds(w http.ResponseWriter, r *http.Request)
	GetItems(w http.ResponseWriter, r *http.Request)
	CreateItems(w http.ResponseWriter, r *http.Request)
	UploadFiles(w http.ResponseWriter, r *http.Request)
	GetRssRedirect(w http.ResponseWriter, r *http.Request)
}

type handlerImpl struct {
	c Controller
}

func NewHandler(c Controller) Handler {
	return &handlerImpl{c}
}

func (h *handlerImpl) GetFeeds(w http.ResponseWriter, r *http.Request) {
	usr := y.Must(user.Get(r))

	req, err := yhttp.DecodeJson[GetFeedsRequest](r.Body)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	response, err := h.c.GetFeeds(r.Context(), &req, usr)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	yhttp.Render(w, r, response).JSON()
}

func (h *handlerImpl) CreateFeeds(w http.ResponseWriter, r *http.Request) {
	usr := y.Must(user.Get(r))

	req, err := yhttp.DecodeJson[CreateFeedsRequest](r.Body)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	response, err := h.c.CreateFeeds(r.Context(), &req, usr)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	yhttp.Render(w, r, response).JSON()
}

func (h *handlerImpl) GetItems(w http.ResponseWriter, r *http.Request) {
	usr := y.Must(user.Get(r))

	req, err := yhttp.DecodeJson[GetItemsRequest](r.Body)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	response, err := h.c.GetItems(r.Context(), &req, usr)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	yhttp.Render(w, r, response).JSON()
}

func (h *handlerImpl) CreateItems(w http.ResponseWriter, r *http.Request) {
	usr := y.Must(user.Get(r))

	req, err := yhttp.DecodeJson[CreateItemsRequest](r.Body)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	response, err := h.c.CreateItems(r.Context(), &req, usr)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	yhttp.Render(w, r, response).JSON()
}

func (h *handlerImpl) UploadFiles(w http.ResponseWriter, r *http.Request) {
	usr := y.Must(user.Get(r))
	ctx := r.Context()

	if err := r.ParseMultipartForm(0); err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	files := make([]multipart.File, 0)
	for _, fileHeader := range r.MultipartForm.File["file"] {
		file, err := fileHeader.Open()
		if err != nil {
			yhttp.Render(w, r, err).JSON()
			return
		}
		files = append(files, file)
	}

	if len(files) == 0 {
		yhttp.Render(w, r, fmt.Errorf("HTTP 400: no files provided")).JSON()
		return
	}

	defer func(fs []multipart.File) {
		logger := slogger.FromContext(ctx)

		for _, f := range fs {
			if err := f.Close(); err != nil {
				logger.Warn("close file", "err", err)
			}
		}
		if err := r.MultipartForm.RemoveAll(); err != nil {
			logger.Warn("remove multipart tmp files", "err", err)
		}
	}(files)

	response, err := h.c.UploadFiles(ctx, files, usr)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	yhttp.Render(w, r, response).JSON()
}

func (h *handlerImpl) GetRssRedirect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "feedId")

	url, err := h.c.GetRssUrl(r.Context(), id)
	if err != nil {
		yhttp.Render(w, r, err).JSON()
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
