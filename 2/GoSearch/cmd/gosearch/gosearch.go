package main

import (
	"flag"
	"fmt"
	"gosearch/2/GoSearch/pkg/crawler/spider"
	"log"
)

func main() {
	var nFlag = flag.String("s", "", "Print links from websites")
	flag.Parse()
	if len(*nFlag) < 1 {
		flag.PrintDefaults()
		return
	}

	urls := [2]string{"https://go.dev", "https://golang.org"}
	depth := 2

	links, err := scanLinks(urls, depth)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	for _, l := range links {
		fmt.Println(l)
	}
}

func scanLinks(urls [2]string, depth int) ([]string, error) {
	spider := spider.New()
	result := []string{}

	for _, u := range urls {
		doc, err := spider.Scan(u, depth)
		if err != nil {
			return nil, err
		}
		for _, d := range doc {
			result = append(result, d.URL)
		}
	}
	return result, nil
}
