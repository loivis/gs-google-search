package gs

import (
	"io"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Result .
type Result struct {
	Link  string
	Title string
	Words string
}

// Search .
func (c *Client) Search(q *url.Values) ([]Result, error) {
	res := []Result{}

	c.searchURL.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, c.searchURL.String(), nil)
	if err != nil {
		return res, err
	}

	req.Header.Set("User-Agent", c.userAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	res = parseBody(resp.Body)

	return res, nil
}

// Search .
func Search(q *url.Values) ([]Result, error) {
	return defaultClient().Search(q)
}

func parseBody(rc io.ReadCloser) []Result {
	res := []Result{}
	doc, _ := goquery.NewDocumentFromReader(rc)

	doc.Find("div#center_col").Find("div.g").Each(func(i int, sel *goquery.Selection) {
		link, _ := sel.Find("div.r").Find("a").Attr("href")
		title := sel.Find("div.r").Find("h3").Text()
		words := sel.Find("div.s").Find("span.st").Text()

		res = append(res, Result{
			Link:  link,
			Title: title,
			Words: words,
		})
	})

	return res
}
