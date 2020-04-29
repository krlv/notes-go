package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/krlv/goweb/note"
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
	r.StrictSlash(true)

	r.HandleFunc("/", PageHandler)

	r.HandleFunc("/notes", note.GetNotes)
	r.HandleFunc("/notes/{id:[0-9]+}", note.GetNote)

	b := r.PathPrefix("/blog").Subrouter()
	b.HandleFunc("", BlogHandler)
	b.HandleFunc("/{slug}", BlogHandler)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	http.Handle("/", r)
	http.ListenAndServe(Port, nil)
}
