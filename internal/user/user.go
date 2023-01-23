package user

type User interface {
	Id() string
}

type noUser struct{}

func (u *noUser) Id() string {
	return ""
}
