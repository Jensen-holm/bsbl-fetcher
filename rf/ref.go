package rf

import (
	"github.com/Jensen-holm/bsbl-api/user"
)

type Reference struct {
	user    *user.User
	headers *map[string]string
	results map[any]any
	request string
}

func NewRef(usr *user.User) *Reference {
	r := usr.Headers["request"]
	return &Reference{
		headers: &usr.Headers,
		user:    usr,
		request: r,
	}
}

func (rf *Reference) Results() map[any]any {
	return rf.results
}

func (rf *Reference) Headers() *map[string]string {
	return rf.headers
}

func (rf *Reference) Request() string {
	return rf.request
}

func (rf *Reference) User() *user.User {
	return rf.user
}

func (rf *Reference) SetResults(r map[any]any) {
	rf.results = r
}
