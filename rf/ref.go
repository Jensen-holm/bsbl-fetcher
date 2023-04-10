package rf

import "github.com/Jensen-holm/bsbl-api/user"

type Reference struct {
	User    *user.User
	Request string
	Headers *map[string]string
	Results string
}

func NewRef(usr *user.User) *Reference {
	r := usr.Headers["request"]
	return &Reference{
		Headers: &usr.Headers,
		User:    usr,
		Request: r,
	}
}
