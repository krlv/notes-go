package main

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/krlv/goweb/blog"
	"github.com/krlv/goweb/note"
)

const (
	Port = ":8080"
)

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

	r.HandleFunc("/blog", blog.GetPages)
	r.HandleFunc("/blog/{slug}", blog.GetPageBySlug)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	http.Handle("/", r)
	http.ListenAndServe(Port, nil)
}
