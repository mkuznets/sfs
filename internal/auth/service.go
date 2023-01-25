package auth

import "net/http"

type Service interface {
	Middleware() func(next http.Handler) http.Handler
}

type NoAuth struct{}

func (*NoAuth) Middleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
