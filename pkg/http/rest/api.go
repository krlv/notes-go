package rest

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/krlv/notes-go/pkg/storage"
)

// DB is a page repository
var db *storage.MemoryStorage

func init() {
	db = storage.NewStorage()
}

// GetNotes returns list of available notes
func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(db.FindNotes())
}

// GetNote returns a single note by note ID
func GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// TODO handle
	}

	n, err := db.GetNoteByID(id)
	if err != nil {
		// TODO handle
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(n)
}

// CreateNote creates new note from submitted data and redirect to the note view page
func CreateNote(w http.ResponseWriter, r *http.Request) {
	var p map[string]interface{}

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		// TODO handle
	}

	// TODO validate payload body

	id, err := db.AddNote(p["title"].(string), p["body"].(string))
	if err != nil {
		// TODO handle
	}

	n, err := db.GetNoteByID(id)
	if err != nil {
		// TODO handle
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(n)
}

// UpdateNote updates note with submitted data and redirect to the note view page
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// TODO handle
	}

	var p map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		// TODO handle
	}

	// TODO validate payload body

	err = db.UpdateNote(id, p["title"].(string), p["body"].(string))
	if err != nil {
		// TODO handle
	}

	n, err := db.GetNoteByID(id)
	if err != nil {
		// TODO handle
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(n)
}

// DeleteNote deletes note and redirects to the note list page
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		// TODO handle
	}

	err = db.DeleteNote(id)
	if err != nil {
		// TODO show error page
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
