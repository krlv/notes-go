package blog

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

// DB is a page repository
var storage Storage

func init() {
	storage = NewStorage()
}

// GetPage renders list of pages from data storage
func GetPages(w http.ResponseWriter, r *http.Request) {
	pages := storage.FindPages()

	t, err := template.ParseFiles("templates/blog.html")
	if err != nil {

	}

	t.Execute(w, pages)
}

// GetPageBySlug returns a single page by it's slug
func GetPageBySlug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	p, err := storage.GetPageBySlug(slug)
	if err != nil {

	}

	t, err := template.ParseFiles("templates/post.html")
	if err != nil {

	}

	t.Execute(w, p)
}
