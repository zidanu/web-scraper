package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(inputURL string) (string, error) {
	urlStruct, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	s := (fmt.Sprintf("%s%s", urlStruct.Host, urlStruct.Path))
	length := len(s)
	if s[length-1] == '/' {
		s = s[:length-1]
	}

	return s, nil
}
