package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	urlStruct, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	if (urlStruct.Port() == "443" && urlStruct.Scheme == "https") || (urlStruct.Port() == "80" && urlStruct.Scheme == "http") {
		urlStruct.Host = urlStruct.Hostname()
	}
	normalizedURL := urlStruct.Host + urlStruct.Path
	if query := urlStruct.RawQuery; query != "" {
		normalizedURL += ("?" + query)
	}
	if normalizedURL != "" {
		normalizedURL = strings.TrimRight(normalizedURL, "/")
	}

	return normalizedURL, nil
}
