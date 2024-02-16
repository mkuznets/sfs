package auth

import "net/http"

type Service interface {
	Middleware(handleError func(w http.ResponseWriter, r *http.Request, err error)) func(next http.Handler) http.Handler
}
