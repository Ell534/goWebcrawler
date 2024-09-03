package main

import (
	//"errors"
	"fmt"
	"golang.org/x/net/html"
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

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing rawBaseURL: %w", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return []string{}, fmt.Errorf("error parsing htmlBody into parse tree: %w", err)
	}

	var urls []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						fmt.Printf("could not parse href '%v': %v\n", a.Val, err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return urls, nil
}
