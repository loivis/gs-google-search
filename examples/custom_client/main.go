package main

import (
	"fmt"
	"net/url"

	gs "github.com/loivis/gs-google-search"
)

func main() {
	c := gs.NewClient()

	query := &url.Values{
		"q":   {"自定义"},
		"num": {"5"},
		"hl":  {"zh-CN"},
	}
	res, err := c.Search(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range res {
		fmt.Printf("%+v\n", r)
	}
}
