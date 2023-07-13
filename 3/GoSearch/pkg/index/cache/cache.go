// Package cache реализует обратный индекс заголовков документов
// Позволяет добавить документ в индекс и получить номера по искомому
// слову в заголовке

package cache

import (
	"gosearch/3/GoSearch/pkg/crawler"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

type Index struct {
	store map[string][]int
}

func New() *Index {
	i := Index{store: map[string][]int{}}
	return &i
}

// Add - фомирует обратный индекс заголовков документов
func (i *Index) Add(docs []crawler.Document) {
	for _, doc := range docs {
		for _, word := range splitWords(doc.Title) {
			if !slices.Contains(i.store[word], doc.ID) {
				i.store[word] = append(i.store[word], doc.ID)
				sort.Ints(i.store[word])
			}
		}
	}
}

// Search - осуществляет поиск по поиск в индексе по искомому слову
func (i *Index) Search(query string) []int {
	return i.store[strings.ToLower(query)]
}

// splitWords - преобразуте строку в массив слов в нижнем регистре
func splitWords(s string) []string {
	var words []string

	for _, word := range strings.Split(s, " ") {
		words = append(words, strings.ToLower(word))
	}

	return words
}
