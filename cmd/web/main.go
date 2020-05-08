package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlv/goweb/pkg/http/web"
)

const (
	Port = ":8080"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	r.HandleFunc("/", web.StaticPage)

	r.HandleFunc("/notes", web.GetNotes).Methods("GET")
	r.HandleFunc("/notes", web.CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id:[0-9]+}", web.GetNote).Methods("GET")
	r.HandleFunc("/notes/{id:[0-9]+}", web.UpdateNote).Methods("PUT")

	r.HandleFunc("/blog", web.GetPages)
	r.HandleFunc("/blog/{slug}", web.GetPageBySlug)

	r.NotFoundHandler = http.HandlerFunc(web.NotFound)

	http.Handle("/", r)
	http.ListenAndServe(Port, nil)
}
