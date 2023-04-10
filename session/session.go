package session

import (
	"net/http"
)

type Session interface {
	SendGet(url string, h map[string]string, client *http.Client)
	ValidRequest(request *map[string]string)
	Main() interface{}
}
