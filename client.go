package gs

import (
	"net/http"
	"net/url"
)

var (
	defaultSearchURL = &url.URL{
		Scheme: "https",
		Host:   "www.google.com",
		Path:   "/search",
	}
	defaultUserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36"
)

// Client .
type Client struct {
	httpClient *http.Client
	searchURL  *url.URL
	userAgent  string
}

func defaultClient() *Client {
	return &Client{
		httpClient: http.DefaultClient,
		searchURL:  defaultSearchURL,
		userAgent:  defaultUserAgent,
	}
}

// NewClient .
func NewClient(opts ...func(*Client)) *Client {
	c := defaultClient()

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithHTTPClient .
func WithHTTPClient(hc *http.Client) func(c *Client) {
	return func(c *Client) {
		c.httpClient = hc
	}
}

// WithSearchURL .
func WithSearchURL(url *url.URL) func(c *Client) {
	return func(c *Client) {
		c.searchURL = url
	}
}

// WithUserAgent .
func WithUserAgent(ua string) func(c *Client) {
	return func(c *Client) {
		c.userAgent = ua
	}
}
