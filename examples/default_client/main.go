package main

import (
	"fmt"
	"net/url"

	gs "github.com/loivis/gs-google-search"
)

func main() {
	query := &url.Values{
		"q":   {"默认"},
		"num": {"5"},
		"hl":  {"zh-CN"},
	}
	res, err := gs.Search(query)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range res {
		fmt.Printf("%+v\n", r)
	}
}
