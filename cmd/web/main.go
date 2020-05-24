package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlv/goweb/pkg/http/rest"
	"github.com/krlv/goweb/pkg/http/web"
)

const (
	Port = ":8080"
)

func main() {
	r := mux.NewRouter()
	r.StrictSlash(true)

	s := r.PathPrefix("/api/notes").Subrouter()
	s.HandleFunc("/", rest.GetNotes).Methods("GET")
	s.HandleFunc("/", rest.CreateNote).Methods("POST")
	s.HandleFunc("/{id:[0-9]+}", rest.GetNote).Methods("GET")
	s.HandleFunc("/{id:[0-9]+}", rest.UpdateNote).Methods("PUT")
	s.HandleFunc("/{id:[0-9]+}", rest.DeleteNote).Methods("DELETE")

	r.HandleFunc("/", web.StaticPage)

	r.HandleFunc("/notes", web.GetNotes).Methods("GET")
	r.HandleFunc("/notes", web.CreateNote).Methods("POST")
	r.HandleFunc("/notes/{id:[0-9]+}", web.GetNote).Methods("GET")
	r.HandleFunc("/notes/{id:[0-9]+}", web.UpdateNote).Methods("PUT")
	r.HandleFunc("/notes/{id:[0-9]+}", web.DeleteNote).Methods("DELETE")

	r.HandleFunc("/blog", web.GetPages)
	r.HandleFunc("/blog/{slug}", web.GetPageBySlug)

	r.NotFoundHandler = http.HandlerFunc(web.NotFound)

	log.Fatal(http.ListenAndServe(Port, r))
}
