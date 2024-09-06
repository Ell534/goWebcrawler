package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) map[string]int {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Fatal(err)
	}
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname() {
		return pages
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := pages[normalizedCurrentURL]; ok {
		pages[normalizedCurrentURL]++
		return pages
	}

	pages[normalizedCurrentURL] = 1

	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("This is the current html from the getHTML function:\n %s", currentHTML)

	nextURLs, err := getURLsFromHTML(currentHTML, rawBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	for i, nextURL := range nextURLs {
		fmt.Printf("Current call to crawlPage: %d\n", i)
		crawlPage(rawBaseURL, nextURL, pages)
	}

	return pages
}
