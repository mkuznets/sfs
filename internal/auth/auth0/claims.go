package auth0

type User struct {
	id string
}

func (c *User) Id() string {
	return c.id
}
