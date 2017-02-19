package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

const (
	Port = ":8080"
)

func PageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug, ok := vars["slug"]

	if !ok {
		slug = "home"
	}

	fileName := "public/" + slug + ".html"
	log.Print("PageHandler:", fileName)

	http.ServeFile(w, r, fileName)
}

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "public/blog.html"
	log.Print("BlogHandler:", fileName)

	http.ServeFile(w, r, fileName)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", PageHandler)
	r.HandleFunc("/{slug}", PageHandler)

	b := r.PathPrefix("/blog").Subrouter()
	b.HandleFunc("/", BlogHandler)
	b.HandleFunc("/{slug}", BlogHandler)

	http.Handle("/", r)
	http.ListenAndServe(Port,nil)
}