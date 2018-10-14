package gs

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
)

type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestClient_Search(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		file, err := os.Open("./testdata/search.result")
		if err != nil {
			t.Fatalf(err.Error())
		}

		c := NewClient(WithHTTPClient(&http.Client{
			Transport: roundTripper(func(*http.Request) (*http.Response, error) {
				return &http.Response{
					Body:       ioutil.NopCloser(file),
					StatusCode: http.StatusInternalServerError,
				}, nil
			}),
		}))

		gotResults, err := c.Search(&url.Values{})
		if err != nil {
			t.Fatalf("err is not nil")
		}

		wantResults := []Result{
			{
				Link:  "https://www.thesaurus.com/browse/example",
				Title: "Example Synonyms, Example Antonyms | Thesaurus.com",
			},
			{
				Link:  "https://www.trythisforexample.com/",
				Title: "Example",
			},
			{
				Link:  "https://sv.wikipedia.org/wiki/Example",
				Title: "Example – Wikipedia",
			},
			{
				Link:  "https://en.wikipedia.org/wiki/Example_(musician)",
				Title: "Example (musician) - Wikipedia",
			},
		}

		if got, want := len(gotResults), len(wantResults); got != want {
			t.Fatalf("got results = %+v", gotResults)
		}
		for i := range wantResults {
			if gotResults[i].Link != wantResults[i].Link {
				t.Errorf("gotResults[%d] = %+v, want %+v", i, gotResults[i], wantResults[i])
			}
		}
	})

	t.Run("HTTPRequestError", func(t *testing.T) {

		c := NewClient(WithHTTPClient(&http.Client{
			Transport: roundTripper(func(*http.Request) (*http.Response, error) {
				return &http.Response{}, errors.New("mock error")
			}),
		}))

		_, err := c.Search(&url.Values{})

		if err == nil {
			t.Fatalf("err is nil")
		}

		if got, want := err.Error(), "mock error"; !strings.Contains(got, want) {
			t.Fatalf("got error = %q, want %q", got, want)
		}
	})
}

func TestParseBody(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		file, err := os.Open("./testdata/search.result")
		if err != nil {
			t.Fatalf(err.Error())
		}

		gotResults := parseBody(file)

		wantResults := []Result{
			{
				Link:  "https://www.thesaurus.com/browse/example",
				Title: "Example Synonyms, Example Antonyms | Thesaurus.com",
			},
			{
				Link:  "https://www.trythisforexample.com/",
				Title: "Example",
			},
			{
				Link:  "https://sv.wikipedia.org/wiki/Example",
				Title: "Example – Wikipedia",
			},
			{
				Link:  "https://en.wikipedia.org/wiki/Example_(musician)",
				Title: "Example (musician) - Wikipedia",
			},
		}

		if got, want := len(gotResults), len(wantResults); got != want {
			t.Fatalf("got results = %+v", gotResults)
		}
		for i := range wantResults {
			if gotResults[i].Link != wantResults[i].Link {
				t.Errorf("gotResults[%d] = %+v, want %+v", i, gotResults[i], wantResults[i])
			}
		}
	})
}
