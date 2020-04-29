package blog

import (
	"strconv"

	"github.com/icrowley/fake"
)

// Storage is an adapter for persistence implementation (port interface)
type Storage interface {
	// GetPageBySlug returns page by it's slug or error if page not found
	FindPages() []Page

	// GetPageBySlug returns page by it's slug or error if page not found
	GetPageBySlug(string) (Page, error)
}

// NewStorage creates new repository
func NewStorage() Storage {
	maxPages := 3

	pages := make(map[string]Page, maxPages)
	for i := 0; i < maxPages; i++ {
		slug := "slug-" + strconv.Itoa(i+1)

		pages[slug] = Page{
			ID:    i + 1,
			Slug:  slug,
			Title: fake.Sentence(),
			Body:  fake.Paragraphs(),
		}
	}

	return MemoryStorage{
		pages: pages,
	}
}

// MemoryStorage repository
type MemoryStorage struct {
	pages map[string]Page
}

// GetPageBySlug returns page by it's slug
func (s MemoryStorage) FindPages() []Page {
	pages := make([]Page, 0, len(s.pages))

	for _, page := range s.pages {
		pages = append(pages, page)
	}

	return pages
}

// GetPageBySlug returns page by it's slug
func (s MemoryStorage) GetPageBySlug(slug string) (Page, error) {
	return s.pages[slug], nil
}
