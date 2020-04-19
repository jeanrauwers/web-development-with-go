package main

import (
	"net/http"

	"gocourse.com/views"

	"github.com/gorilla/mux"
)

var (
	homeView     *views.View
	contactView  *views.View
	notFoundView *views.View
)

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	contactView.Render(w, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	homeView.Render(w, nil)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	notFoundView.Render(w, nil)
}

func main() {

	homeView = views.NewView("main", "views/home.gohtml")
	contactView = views.NewView("main", "views/contact.gohtml")
	notFoundView = views.NewView("main", "views/notFound.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = r.NewRoute().HandlerFunc(notFoundHandler).GetHandler()

	http.ListenAndServe("localhost:3333", r)
}
