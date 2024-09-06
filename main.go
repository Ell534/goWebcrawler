package main

import (
	"fmt"
	// "log"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	rawBaseURL := os.Args[1]

	fmt.Printf("starting crawl of: %s\n", rawBaseURL)

	// html, err := getHTML(rawBaseURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(html)

	pagesMap := crawlPage(rawBaseURL, rawBaseURL, map[string]int{})

	for key, value := range pagesMap {
		fmt.Printf("Key: %s\nValue: %d\n", key, value)
	}
}
