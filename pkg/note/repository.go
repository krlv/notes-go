package note

// Repository is an adapter for persistence implementation (port interface)
type Repository interface {
	// FindNotes returns list of existing notes
	FindNotes() []Note

	// GetNoteByID returns note by ID or error if note not found
	GetNoteByID(int) (Note, error)
}
