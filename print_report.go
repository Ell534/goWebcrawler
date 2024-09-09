package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type pagesStruct struct {
	pageURL   string
	pageCount int
}

func sortPages(pages map[string]int) []pagesStruct {
	sortedPages := []pagesStruct{}
	for url, hits := range pages {
		sortedPages = append(sortedPages, pagesStruct{pageURL: url, pageCount: hits})
	}
	slices.SortFunc(sortedPages, func(a, b pagesStruct) int {
		if n := cmp.Compare(b.pageCount, a.pageCount); n != 0 {
			return n
		}
		return strings.Compare(a.pageURL, b.pageURL)
	})
	return sortedPages
}

func printReport(pages map[string]int, baseURL string) {
	pagesToPrint := sortPages(pages)
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)
	for _, page := range pagesToPrint {
		fmt.Printf("Found %d internal links to %s\n", page.pageCount, page.pageURL)
	}
}
