package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if len(cfg.pages) >= cfg.maxPages {
		return
	}

	if _, ok := cfg.pages[normalizedURL]; ok {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing rawCurrentURL: '%s', error: %w\n", rawCurrentURL, err)
	}

	if cfg.baseURL.Hostname() != parsedCurrentURL.Hostname() {
		return
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing url: %w\n", err)
	}

	isFirst := cfg.addPageVisit(normalizedCurrentURL)
	if !isFirst {
		return
	}

	currentHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error from getHTML: %w\n", err)
	}
	// fmt.Printf("This is the current html from the getHTML function:\n %s", currentHTML)

	nextURLs, err := getURLsFromHTML(currentHTML, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting URLs from HTML: %w\n", err)
	}

	for _, nextURL := range nextURLs {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
