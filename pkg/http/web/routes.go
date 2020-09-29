package web

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/krlv/goweb/pkg/storage"
)

// DB is a page repository
var db *storage.MemoryStorage

func init() {
	db = storage.NewStorage()
}

// NotFound will be invoked by mux when URL can't be matched
// to registered routes (404 page)
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Print("NotFoundHandler:", r.RequestURI)

	path := strings.TrimLeft(r.RequestURI, "/")
	err := errors.New("handler not found for path " + path)

	notFoundPage(err, w, r)
}

// StaticPage renders static pages
func StaticPage(w http.ResponseWriter, r *http.Request) {
	path, _ := mux.CurrentRoute(r).GetPathTemplate()
	path = strings.TrimLeft(path, "/")

	if path == "" {
		path = "home"
	}

	fileName := "web/static/" + path + ".html"

	http.ServeFile(w, r, fileName)
}

// GetPage renders list of pages from data storage
func GetPages(w http.ResponseWriter, r *http.Request) {
	pages := db.FindPages()

	t, err := template.ParseFiles("web/template/blog.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, pages)
	if err != nil {
		panic(err)
	}
}

// GetPageBySlug returns a single page by it's slug
func GetPageBySlug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	p, err := db.GetPageBySlug(slug)
	if err != nil {
		notFoundPage(err, w, r)
		return
	}

	t, err := template.ParseFiles("web/template/post.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

// GetNotes returns list of notes from data storage
func GetNotes(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("web/template/notes.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, db.FindNotes())
	if err != nil {
		panic(err)
	}
}

// CreateNote creates new note from submitted data and redirect to the note view page
func CreateNote(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	body := r.FormValue("note")

	id, err := db.AddNote(title, body)
	if err != nil {
		errorPage(err, w, r)
		return
	}

	http.Redirect(w, r, "/notes/"+strconv.Itoa(id), http.StatusFound)
}

// GetNote returns a single note by note ID
func GetNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid note ID.", http.StatusBadRequest)
		return
	}

	n, err := db.GetNoteByID(id)
	if err != nil {
		notFoundPage(err, w, r)
		return
	}

	t, err := template.ParseFiles("web/template/note.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, n)
	if err != nil {
		panic(err)
	}
}

// UpdateNote updates note with submitted data and redirect to the note view page
func UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid note ID.", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	body := r.FormValue("note")

	err = db.UpdateNote(id, title, body)
	if err != nil {
		notFoundPage(err, w, r)
		return
	}

	http.Redirect(w, r, "/notes/"+strconv.Itoa(id), http.StatusFound)
}

// DeleteNote deletes note and redirects to the note list page
func DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid note ID.", http.StatusBadRequest)
		return
	}

	err = db.DeleteNote(id)
	if err != nil {
		notFoundPage(err, w, r)
		return
	}

	http.Redirect(w, r, "/notes", http.StatusFound)
}
