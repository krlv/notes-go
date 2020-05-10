package rest

import (
	"encoding/json"
	"net/http"
)

// CreateNote creates new note from submitted data and redirect to the note view page
func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// GetNote returns a single note by note ID
func GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// GetNotes returns list of available notes
func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// UpdateNote updates note with submitted data and redirect to the note view page
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

// DeleteNote deletes note and redirects to the note list page
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
