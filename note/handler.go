package note

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DB is a notes repository
var storage Storage

func init() {
	storage = NewStorage()
}

// GetNotes returns list of notes from data storage
func GetNotes(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/notes.html")
	if err != nil {

	}

	err = t.Execute(w, storage.FindNotes())
	if err != nil {

	}
}

// GetNote returns a single note by note ID
func GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {

	}

	n, err := storage.GetNoteByID(id)
	if err != nil {

	}

	t, err := template.ParseFiles("templates/note.html")
	if err != nil {

	}

	t.Execute(w, n)
}
