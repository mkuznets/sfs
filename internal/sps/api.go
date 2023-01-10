package sps

import (
	"github.com/go-chi/render"
	"net/http"
)

type Api struct {
}

func (a *Api) ListChannels(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{"hello": "world"})
	return
}
