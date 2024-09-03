package main

import (
	"fmt"
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

	fmt.Printf("starting crawl of: %s", rawBaseURL)

	// if len(os.Args) == 2 {
	// 	fmt.Printf("starting crawl of: %s", os.Args[1])
	// }
}
