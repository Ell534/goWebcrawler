package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("error executing http GET request to %s: %w", rawURL, err)
	}

	if res.StatusCode > 399 {
		return "", fmt.Errorf("error level status code from http response: %d", res.StatusCode)
	}

	if res.Header.Get("content-type") != "text/html" {
		return "", fmt.Errorf("content-type header is not text/html it is: %s", res.Header.Get("content-type"))
	}

	html, err := io.ReadAll(res.Body)

	defer res.Body.Close()
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	htmlString := string(html)

	return htmlString, nil
}
