package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(fullURL string) (string, error) {
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", fmt.Errorf("error parsing given URL: %w", err)
	}
	normalizedURL := parsedURL.Hostname() + parsedURL.Path
	normalizedURL = strings.ToLower(normalizedURL)
	normalizedURL = strings.TrimSuffix(normalizedURL, "/")
	return normalizedURL, nil
}
