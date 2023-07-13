package index

import (
	"gosearch/3/GoSearch/pkg/crawler"
)

type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
