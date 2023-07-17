package main

import (
	"flag"
	"fmt"
	"gosearch/3/GoSearch/pkg/crawler"
	"gosearch/3/GoSearch/pkg/crawler/spider"
	"gosearch/3/GoSearch/pkg/index"
	"gosearch/3/GoSearch/pkg/index/cache"
	"gosearch/3/GoSearch/pkg/storage"
	"gosearch/3/GoSearch/pkg/storage/memstore"
	"log"
	"math/rand"
	"time"
)

type searcher struct {
	storage storage.Interface
	index   index.Interface
	sites   []string
	depth   int
}

func main() {
	app := new()
	app.run()
}

func new() *searcher {
	searcher := searcher{}
	searcher.storage = memstore.New()
	searcher.index = cache.New()
	searcher.sites = []string{"https://go.dev", "https://golang.org"}
	searcher.depth = 2

	return &searcher
}

func (s *searcher) run() {
	var query = flag.String("s", "", "Plese enter word you want to search:")
	flag.Parse()
	if *query == "" {
		flag.PrintDefaults()
		return
	}
	fmt.Printf("Please wait, searching docs with word - '%s'\n", *query)
	rand.Seed(time.Now().UnixNano())
	spider := spider.New()

	for _, url := range s.sites {
		docs, errs := spider.Scan(url, s.depth)
		if errs != nil {
			log.Println(errs)
			continue
		}

		for _, doc := range docs {
			doc.ID = rand.Intn(10000)
			s.storage.Add([]crawler.Document{doc})
			s.index.Add([]crawler.Document{doc})
		}
	}

	ids := s.index.Search(*query)
	fmt.Println("Index ids: ", ids)

	fmt.Println("Storage docs:")
	for _, doc := range s.storage.Search(ids) {
		fmt.Println("- ", doc)
	}
}
