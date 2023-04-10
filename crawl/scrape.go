package crawl

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func SendGet(
	url string,
	headers *map[string]string,
	client *http.Client,
) (*goquery.Document, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	addHeaders(*headers, req)
	doc, err := readBody(client, req)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func readBody(client *http.Client, req *http.Request) (*goquery.Document, error) {
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func addHeaders(headers map[string]string, req *http.Request) {
	for key, value := range headers {
		req.Header.Add(key, value)
	}
}

func GetHeaders(c *fiber.Ctx) map[string]string {
	h := make(map[string]string)

	c.Request().Header.VisitAll(func(key, value []byte) {
		h[string(key)] = string(value)
	})
	return h
}
