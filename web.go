package main

import (
	"github.com/gorilla/mux"
	"github.com/icrowley/fake"
	"github.com/krlv/goweb/entity"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"text/template"
)

const (
	Port = ":8080"
)

func getRoutePath(r *http.Request) (*url.URL, error) {
	var pairs []string

	vars := mux.Vars(r)
	if len(vars) > 0 {
		pairs = make([]string, len(vars)*2)
		for key, value := range vars {
			pairs = append(pairs, key, value)
		}
	}

	return mux.CurrentRoute(r).URLPath(pairs...)
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := mux.CurrentRoute(r).GetPathTemplate()
	path = strings.TrimLeft(path, "/")

	if path == "" {
		path = "home"
	}

	fileName := "public/" + path + ".html"
	log.Print("PageHandler:", fileName)

	http.ServeFile(w, r, fileName)
}

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := getRoutePath(r)
	log.Print("NotesHandler:", path)

	notes := make([]entity.Note, 2)
	for i := 0; i < cap(notes); i++ {
		note := new(entity.Note)
		note.Title = fake.Sentence()
		note.Body = fake.Paragraphs()

		note.Tags = make([]string, rand.Intn(3))
		for j := 0; j < cap(note.Tags); j++ {
			note.Tags[j] = fake.Word()
		}

		notes[i] = *note
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	if id > 0 {
		t, _ := template.ParseFiles("templates/note.html")
		t.Execute(w, notes[id])
	} else {
		t, _ := template.ParseFiles("templates/notes.html")
		t.Execute(w, notes)
	}
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	path, _ := getRoutePath(r)
	log.Print("BlogHandler:", path)

	slug := mux.Vars(r)["slug"]
	slug = strings.Title(slug)

	t, _ := template.ParseFiles("templates/blog.html")
	t.Execute(w, map[string]interface{}{"Title": slug})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("NotFoundHandler:", r.RequestURI)
	path := strings.TrimLeft(r.RequestURI, "/")

	t, _ := template.ParseFiles("templates/404.html")
	t.Execute(w, map[string]interface{}{"Path": path})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", PageHandler)

	n := r.PathPrefix("/notes").Subrouter()
	n.HandleFunc("/", NotesHandler)
	n.HandleFunc("/{id:[0-9]+}", NotesHandler)

	b := r.PathPrefix("/blog").Subrouter()
	b.HandleFunc("", BlogHandler)
	b.HandleFunc("/{slug}", BlogHandler)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	http.Handle("/", r)
	http.ListenAndServe(Port, nil)
}
