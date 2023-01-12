package auth

import (
	"net/http"
)

func FakeAuth(userId string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user := &userImpl{
				id: userId,
			}
			r = withUser(r, user)
			next.ServeHTTP(w, r)
		})
	}
}
