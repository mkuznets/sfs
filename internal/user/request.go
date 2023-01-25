package user

import (
	"context"
	"mkuznets.com/go/sfs/ytils/yerr"
	"net/http"
)

type contextKey int

var ctxUserKey = contextKey(0x55)

func Get(r *http.Request) (User, error) {
	user := r.Context().Value(ctxUserKey)
	if user == nil {
		return (*noUser)(nil), nil
	}
	if u, ok := user.(User); ok {
		return u, nil
	} else {
		return nil, yerr.New("invalid user type")
	}
}

func Ctx(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, ctxUserKey, user)
}
