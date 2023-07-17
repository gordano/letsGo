package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler"
	"gosearch/pkg/crawler/spider"
	"log"
	"strings"
)

func main() {
	var query = flag.String("s", "", "Print links from websites")
	flag.Parse()
	if len(*query) < 1 {
		flag.PrintDefaults()
		return
	}

	urls := [2]string{"https://go.dev", "https://golang.org"}
	depth := 2
	docs := scanLinks(urls, depth)
	search(docs, *query)
}

func scanLinks(urls [2]string, depth int) []crawler.Document {
	spider := spider.New()
	result := []crawler.Document{}

	for _, u := range urls {
		doc, err := spider.Scan(u, depth)
		if err != nil {
			log.Print(err)
			continue
		}
		result = append(result, doc...)
	}
	return result
}

func search(docs []crawler.Document, query string) {
	fmt.Println("Wait... your search word is: ", query)
	for _, d := range docs {
		if strings.Contains(d.URL, query) {
			fmt.Printf("Title: %v - Link: %v\n", d.Title, d.URL)
		}
	}
}
