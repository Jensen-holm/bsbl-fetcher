package rf

import (
	"github.com/Jensen-holm/bsbl-api/session"
	"net/http"
)

const BASEURL string = "https://baseball-reference.com"

func (rf *Reference) Main() error {
	err := ValidRequest(rf.Request)
	if err != nil {
		return err
	}

	r, err := session.SendGet(BASEURL, rf.User.Headers, &http.Client{})
	if err != nil {
		return err
	}

	rf.Results = string(r)
	return nil
}

func ValidRequest(r string) error {
	return nil
}
