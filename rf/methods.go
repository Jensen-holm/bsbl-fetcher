package rf

import (
	scrape "github.com/Jensen-holm/bsbl-api/crawl"
	"net/http"
)

const BASEURL string = "https://baseball-reference.com"

func (rf *Reference) Main() error {
	err := rf.ValidRequest(rf.Request())
	if err != nil {
		return err
	}

	r, err := scrape.SendGet(
		BASEURL,
		rf.Headers(),
		&http.Client{},
	)
	if err != nil {
		return err
	}

	rf.SetResults(r.Text())
	return nil
}

func (rf *Reference) ValidRequest(r string) error {
	return nil
}
