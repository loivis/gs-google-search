package gs

import (
	"net/http"
	"net/url"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	t.Run("DefaultConfig", func(t *testing.T) {
		c := NewClient()

		if got, want := c.httpClient, http.DefaultClient; got != want {
			t.Errorf("c.httpClient = %v, want %v", got, want)
		}
		if got, want := c.searchURL.String(), defaultSearchURL.String(); got != want {
			t.Errorf("c.searchURL = %q, want %q", got, want)
		}
		if got, want := c.userAgent, defaultUserAgent; got != want {
			t.Errorf("c.userAgent = %v, want %v", got, want)
		}
	})

	t.Run("WithHTTPClient", func(t *testing.T) {
		hc := &http.Client{Timeout: time.Duration(2) * time.Second}
		c := NewClient(WithHTTPClient(hc))

		if got, want := c.httpClient, hc; got != want {
			t.Errorf("c.httpClient = %+v, want %+v", got, want)
		}
	})

	t.Run("WithSearchURL", func(t *testing.T) {
		url := &url.URL{Scheme: "https", Host: "www.example.com"}
		c := NewClient(WithSearchURL(url))

		if got, want := c.searchURL.String(), url.String(); got != want {
			t.Errorf("c.searchURL = %+v, want %+v", got, want)
		}
	})

	t.Run("WithUserAgent", func(t *testing.T) {
		ua := "custom user agent"
		c := NewClient(WithUserAgent(ua))

		if got, want := c.userAgent, ua; got != want {
			t.Errorf("c.userAgent = %+v, want %+v", got, want)
		}
	})

}
