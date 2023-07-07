package main

import (
	"flag"
	"fmt"
	"gosearch/pkg/crawler/spider"
)

type scanData struct {
	sites []string
	urls  []string
	depth int
}

func (s *scanData) scan() {
	spider := spider.New()

	for _, site := range s.sites {
		data, er := spider.Scan(site, s.depth)

		if er != nil {
			fmt.Println("Error: ", er)
		}

		for _, e := range data {
			s.urls = append(s.urls, e.URL)
		}
	}
}

func (s *scanData) showLinks() {
	for _, url := range s.urls {
		fmt.Println(url)
	}
}

func initial() bool {
	var nFlag = flag.String("s", "", "Print links from websites")
	flag.Parse()

	return len(*nFlag) > 1
}

func main() {

	if !initial() {
		fmt.Println("Please retry with [-s document]")
		return
	}

	var webData = &scanData{
		sites: []string{"https://go.dev", "https://golang.org"},
		depth: 1,
	}

	webData.scan()
	webData.showLinks()
}
