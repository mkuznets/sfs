package auth

import (
	"context"
	"fmt"
	"mkuznets.com/go/sps/internal/ytils/yerr"
	"net/http"
)

type contextKey string

var ctxUserKey = contextKey("User")

func GetUser(r *http.Request) (User, error) {
	user := r.Context().Value(ctxUserKey)
	if user == nil {
		return nil, yerr.Unauthorised("unauthorised")
	}
	if u, ok := user.(User); ok {
		return u, nil
	} else {
		return nil, fmt.Errorf("invalid user type")
	}
}

func withUser(r *http.Request, user User) *http.Request {
	ctx := context.WithValue(r.Context(), ctxUserKey, user)
	return r.WithContext(ctx)
}
