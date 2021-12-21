package storage

import (
	"math/rand"
	"strconv"
	"sync"

	"github.com/icrowley/fake"
	"github.com/krlv/notes-go/pkg/blog"
	"github.com/krlv/notes-go/pkg/note"
)

// MemoryStorage repository
type MemoryStorage struct {
	sync.RWMutex
	pages map[string]*blog.Page
	notes map[int]*note.Note
}

// NewStorage creates new repository
func NewStorage() *MemoryStorage {
	maxPages := 3

	pages := make(map[string]*blog.Page)
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
	notes := make(map[int]*note.Note)
	for i := 0; i < maxNotes; i++ {
		tags := make([]string, rand.Intn(maxTags))
		for j := 0; j < cap(tags); j++ {
			tags[j] = fake.Word()
		}

		notes[i+1] = &note.Note{
			ID:    i + 1,
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

	for _, p := range s.pages {
		pages = append(pages, p)
	}

	return pages
}

// GetPageBySlug returns page by it's slug
func (s *MemoryStorage) GetPageBySlug(slug string) (*blog.Page, error) {
	p, ok := s.pages[slug]
	if !ok {
		return nil, blog.ErrNotFound
	}

	return p, nil
}

// FindNotes returns notes from memory repository
func (s *MemoryStorage) FindNotes() []note.Note {
	notes := make([]note.Note, 0, len(s.notes))

	for _, n := range s.notes {
		notes = append(notes, *n)
	}

	return notes
}

// GetNoteByID returns note by ID or error if note not found in memory repo
func (s *MemoryStorage) GetNoteByID(id int) (note.Note, error) {
	var n note.Note

	s.RLock()
	np, ok := s.notes[id]
	s.RUnlock()

	if !ok {
		return n, note.ErrNotFound
	}

	return *np, nil
}

// AddNote creates new note and returns it's ID
func (s *MemoryStorage) AddNote(title, body string) (int, error) {
	s.Lock()

	// TODO create proper ID generation
	id := len(s.notes) + 1

	s.notes[id] = &note.Note{
		ID:    id,
		Title: title,
		Body:  body,
	}

	s.Unlock()

	return id, nil
}

// UpdateNote creates new note and returns it's ID
func (s *MemoryStorage) UpdateNote(id int, title string, body string) error {
	s.Lock()
	defer s.Unlock()

	n, ok := s.notes[id]
	if !ok {
		return note.ErrNotFound
	}

	n.Title = title
	n.Body = body

	return nil
}

// DeleteNote removes note by id
func (s *MemoryStorage) DeleteNote(id int) error {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.notes[id]; !ok {
		return note.ErrNotFound
	}

	delete(s.notes, id)

	return nil
}
