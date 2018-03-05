package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"strings"
	"net/url"
)

const (
	Port = ":8080"
)

func getRoutePath(r *http.Request) (*url.URL, error) {
	var pairs []string

	vars := mux.Vars(r)
	if len(vars) > 0 {
		pairs = make([]string, len(vars) * 2)
		for  key, value := range vars {
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

	fileName := "public/blog.html"
	log.Print("BlogHandler:", path)

	http.ServeFile(w, r, fileName)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fileName := "public/404.html"
	log.Print("NotFoundHandler:", r.RequestURI)

	http.ServeFile(w, r, fileName)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", PageHandler)
	r.HandleFunc("/contacts", PageHandler)

	b := r.PathPrefix("/blog").Subrouter()
	b.HandleFunc("", BlogHandler)
	b.HandleFunc("/{slug}", BlogHandler)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	http.Handle("/", r)
	http.ListenAndServe(Port,nil)
}