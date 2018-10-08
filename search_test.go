package gs

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type transporter func(*http.Request) (*http.Response, error)

func (t transporter) RoundTrip(r *http.Request) (*http.Response, error) {
	return t(r)
}

func TestClient_Search(t *testing.T) {
	t.Run("MockRoundTripper", func(t *testing.T) {
		var mockT transporter = func(*http.Request) (*http.Response, error) {
			return &http.Response{
				Body:       ioutil.NopCloser(strings.NewReader("body content")),
				StatusCode: http.StatusInternalServerError,
			}, errors.New("mock error")
		}
		hc := &http.Client{Transport: mockT}

		c := NewClient(WithHTTPClient(hc))
		q := &url.Values{}

		_, err := c.Search(q)
		if err == nil {
			t.Fatalf("err is nil")
		}
	})
}
