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
		n := new(Note)

		n.Title = fake.Sentence()
		n.Body = fake.Paragraphs()

		n.Tags = make([]string, rand.Intn(maxTags))
		for j := 0; j < cap(n.Tags); j++ {
			n.Tags[j] = fake.Word()
		}

		notes[i] = *n
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
	return []Note{}
}

// GetNoteByID returns note by ID or error if note not found in memory repo
func (s MemoryStorage) GetNoteByID(id int) (Note, error) {
	return s.notes[id], nil
}
