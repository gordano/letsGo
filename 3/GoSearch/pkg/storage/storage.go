package storage

import (
	"gosearch/3/GoSearch/pkg/crawler"
)

type Interface interface {
	Add([]crawler.Document)
	Search([]int) []crawler.Document
}
