package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error executing http GET request to %s: %w", rawURL, err)
	}

	defer res.Body.Close()

	if res.StatusCode > 399 {
		return "", fmt.Errorf("error level status code from http response: %d", res.StatusCode)
	}

	contentType := res.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("content-type header is not text/html it is: %s", res.Header.Get("Content-Type"))
	}

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	htmlString := string(html)

	return htmlString, nil
}
