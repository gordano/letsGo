package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler/spider"
	"log"
)

func main() {
	var nFlag = flag.String("s", "", "Print links from websites")
	flag.Parse()
	if len(*nFlag) < 1 {
		flag.PrintDefaults()
		return
	}

	urls := [2]string{"hhtttps://go.dev", "https://golang.org"}
	depth := 2

	links := scanLinks(urls, depth)

	for _, l := range links {
		fmt.Println(l)
	}
}

func scanLinks(urls [2]string, depth int) []string {
	spider := spider.New()
	result := []string{}

	for _, u := range urls {
		doc, err := spider.Scan(u, depth)
		if err != nil {
			log.Print(err)
			continue
		}
		for _, d := range doc {
			result = append(result, d.URL)
		}
	}
	return result
}
