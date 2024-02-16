package auth

import (
	"context"
)

type userKey struct{}

type User interface {
	ID() string
}

type noUser struct{}

func (u *noUser) ID() string {
	return ""
}

func UserFromContext(ctx context.Context) User {
	if user, ok := ctx.Value(userKey{}).(User); ok {
		return user
	}
	return (*noUser)(nil)
}

func NewUserContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userKey{}, user)
}
