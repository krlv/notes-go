package web

import (
	"html/template"
	"net/http"
)

// errorPage sends 500 error page to response writer
func errorPage(error error, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	t, err := template.ParseFiles("web/template/500.html")
	if err != nil {
		panic(err)
	}

	// TODO: show debug data on dev env
	err = t.Execute(w, map[string]interface{}{"Debug": error.Error()})
	if err != nil {
		panic(err)
	}
}

// notFoundPage sends 404 error page to response writer
func notFoundPage(error error, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	t, err := template.ParseFiles("web/template/404.html")
	if err != nil {
		panic(err)
	}

	// TODO: show debug data on dev env
	err = t.Execute(w, map[string]interface{}{"Debug": error.Error()})
	if err != nil {
		panic(err)
	}
}

