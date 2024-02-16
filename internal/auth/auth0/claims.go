package auth0

type User struct {
	id string
}

func (c *User) ID() string {
	return c.id
}
