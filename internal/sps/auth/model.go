package auth

type User interface {
	Id() string
}

type userImpl struct {
	id string
}

func (u *userImpl) Id() string {
	return u.id
}
