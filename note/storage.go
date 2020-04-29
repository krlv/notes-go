package note

import (
	"math/rand"

	"github.com/icrowley/fake"
)

// Storage is an adapter for persistence implementation (port interface)
type Storage interface {
	// FindNotes returns list of existing notes
	FindNotes() []Note

	// GetNoteByID returns note by ID or error if note not found
	GetNoteByID(int) (Note, error)
}

// NewStorage creates new repository
func NewStorage() Storage {
	maxNotes, maxTags := 2, 3

	notes := make(map[int]Note, maxNotes)
	for i := 0; i < maxNotes; i++ {
		tags := make([]string, rand.Intn(maxTags))
		for j := 0; j < cap(tags); j++ {
			tags[j] = fake.Word()
		}

		notes[i] = Note{
			ID: i,
			Title: fake.Sentence(),
			Body: fake.Paragraphs(),
			Tags: tags,
		}
	}

	return MemoryStorage{
		notes: notes,
	}
}

// MemoryStorage repository
type MemoryStorage struct {
	notes map[int]Note
}

// FindNotes returns notes from memory repository
func (s MemoryStorage) FindNotes() []Note {
	notes := make([]Note, 0, len(s.notes))

	for _, page := range s.notes {
		notes = append(notes, page)
	}

	return notes
}

// GetNoteByID returns note by ID or error if note not found in memory repo
func (s MemoryStorage) GetNoteByID(id int) (Note, error) {
	return s.notes[id], nil
}
