package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	fmt.Println("Hello, World!")

	if len(os.Args) < 4 {
		fmt.Println("format should be './crawler' <URL> <maxConcurrency int> <maxPages int>")
		os.Exit(1)
	}

	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]
	maxConcurrency, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Errorf("Error converting cli argument into int: %w", err)
	}
	maxPages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Errorf("Error converting cli argument into int: %w", err)
	}

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	// html, err := getHTML(rawBaseURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(html)

	parsedRawBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Fatal(err)
	}

	crawlConfig := config{
		pages:              map[string]int{},
		baseURL:            parsedRawBaseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	crawlConfig.wg.Add(1)
	go crawlConfig.crawlPage(rawBaseURL)

	crawlConfig.wg.Wait()

	for key, value := range crawlConfig.pages {
		fmt.Printf("Key: %s\nValue: %d\n", key, value)
	}

	printReport(crawlConfig.pages, rawBaseURL)
}
