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

func main() {
	var nFlag = flag.String("s", "", "Print links from websites")
	flag.Parse()

	if len(*nFlag) < 1 {
		fmt.Println("Please retry with [-s document]")
		return
	}

	var webData = &scanData{
		sites: []string{"https://go.dev", "https://golang.org"},
		depth: 2,
	}

	webData.scan()
	webData.showLinks()
}

func (s *scanData) scan() (*scanData, error) {
	spider := spider.New()

	for _, site := range s.sites {
		data, err := spider.Scan(site, s.depth)
		if err != nil {
			return s, err
		}

		for _, e := range data {
			s.urls = append(s.urls, e.URL)
		}
	}
	return s, nil
}

func (s *scanData) showLinks() *scanData {
	for _, url := range s.urls {
		fmt.Println(url)
	}
	return s
}
