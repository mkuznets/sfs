package user

import (
	"context"
	"fmt"
	"mkuznets.com/go/sfs/internal/ytils/yerr"
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
		return nil, fmt.Errorf("invalid user type")
	}
}

func MustGet(r *http.Request) User {
	user, err := Get(r)
	if err != nil {
		panic(err)
	}
	return user
}

func Require(r *http.Request) (User, error) {
	user, err := Get(r)
	if err != nil || user == nil {
		return nil, yerr.Unauthorised("unauthorised").WithCause(err)
	}
	return user, nil
}

func Ctx(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, ctxUserKey, user)
}
