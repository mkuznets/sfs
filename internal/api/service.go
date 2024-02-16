package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mkuznets.com/go/ytils/yhttp"

	"mkuznets.com/go/sfs/internal/auth"
	"mkuznets.com/go/sfs/internal/render"
	"mkuznets.com/go/sfs/internal/slogger"
)

type Service struct {
	controller Controller
}

func NewService(c Controller) *Service {
	return &Service{controller: c}
}

func (s *Service) GetFeeds(w http.ResponseWriter, r *http.Request) {
	user := auth.UserFromContext(r.Context())

	req, err := yhttp.DecodeJson[GetFeedsRequest](r.Body)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	response, err := s.controller.GetFeeds(r.Context(), &req, user)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	render.New(w, r, response).JSON()
}

func (s *Service) CreateFeeds(w http.ResponseWriter, r *http.Request) {
	user := auth.UserFromContext(r.Context())

	req, err := yhttp.DecodeJson[CreateFeedsRequest](r.Body)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	response, err := s.controller.CreateFeeds(r.Context(), &req, user)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	render.New(w, r, response).JSON()
}

func (s *Service) GetItems(w http.ResponseWriter, r *http.Request) {
	user := auth.UserFromContext(r.Context())

	req, err := yhttp.DecodeJson[GetItemsRequest](r.Body)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	response, err := s.controller.GetItems(r.Context(), &req, user)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	render.New(w, r, response).JSON()
}

func (s *Service) CreateItems(w http.ResponseWriter, r *http.Request) {
	user := auth.UserFromContext(r.Context())

	req, err := yhttp.DecodeJson[CreateItemsRequest](r.Body)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	response, err := s.controller.CreateItems(r.Context(), &req, user)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	render.New(w, r, response).JSON()
}

func (s *Service) UploadFiles(w http.ResponseWriter, r *http.Request) {
	user := auth.UserFromContext(r.Context())
	ctx := r.Context()

	if err := r.ParseMultipartForm(0); err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	files := make([]multipart.File, 0)
	for _, fileHeader := range r.MultipartForm.File["file"] {
		file, err := fileHeader.Open()
		if err != nil {
			slogger.WithError(r.Context(), err)
			render.New(w, r, err).JSON()
			return
		}
		files = append(files, file)
	}

	if len(files) == 0 {
		render.New(w, r, fmt.Errorf("HTTP 400: no files provided")).JSON()
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

	response, err := s.controller.UploadFiles(ctx, files, user)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	render.New(w, r, response).JSON()
}

func (s *Service) GetRssRedirect(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "feedId")

	url, err := s.controller.GetRssUrl(r.Context(), id)
	if err != nil {
		slogger.WithError(r.Context(), err)
		render.New(w, r, err).JSON()
		return
	}

	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
