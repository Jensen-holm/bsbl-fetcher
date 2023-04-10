package crawl

import (
	"io/ioutil"
	"net/http"
)

func SendGet(
	url string,
	headers map[string]string,
	client *http.Client,
) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	addHeaders(headers, req)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func addHeaders(headers map[string]string, req *http.Request) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}
