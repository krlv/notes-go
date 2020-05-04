package storage

import (
	"math/rand"
	"strconv"

	"github.com/icrowley/fake"
	"github.com/krlv/goweb/pkg/blog"
	"github.com/krlv/goweb/pkg/note"
)

// MemoryStorage repository
type MemoryStorage struct {
	pages map[string]*blog.Page
	notes map[int]*note.Note
}

// NewStorage creates new repository
func NewStorage() *MemoryStorage {
	maxPages := 3

	pages := make(map[string]*blog.Page, maxPages)
	for i := 0; i < maxPages; i++ {
		slug := "slug-" + strconv.Itoa(i+1)

		pages[slug] = &blog.Page{
			ID:    i + 1,
			Slug:  slug,
			Title: fake.Sentence(),
			Body:  fake.Paragraphs(),
		}
	}

	maxNotes, maxTags := 2, 3
	notes := make(map[int]*note.Note, maxNotes)
	for i := 0; i < maxNotes; i++ {
		tags := make([]string, rand.Intn(maxTags))
		for j := 0; j < cap(tags); j++ {
			tags[j] = fake.Word()
		}

		notes[i] = &note.Note{
			ID:    i,
			Title: fake.Sentence(),
			Body:  fake.Paragraphs(),
			Tags:  tags,
		}
	}

	return &MemoryStorage{
		pages: pages,
		notes: notes,
	}
}

// GetPageBySlug returns page by it's slug
func (s *MemoryStorage) FindPages() []*blog.Page {
	pages := make([]*blog.Page, 0, len(s.pages))

	for _, page := range s.pages {
		pages = append(pages, page)
	}

	return pages
}

// GetPageBySlug returns page by it's slug
func (s *MemoryStorage) GetPageBySlug(slug string) (*blog.Page, error) {
	return s.pages[slug], nil
}

// FindNotes returns notes from memory repository
func (s *MemoryStorage) FindNotes() []*note.Note {
	notes := make([]*note.Note, 0, len(s.notes))

	for _, page := range s.notes {
		notes = append(notes, page)
	}

	return notes
}

// GetNoteByID returns note by ID or error if note not found in memory repo
func (s *MemoryStorage) GetNoteByID(id int) (*note.Note, error) {
	return s.notes[id], nil
}
