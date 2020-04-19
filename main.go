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
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := notFoundView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func main() {

	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	notFoundView = views.NewView("views/notFound.gohtml")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = r.NewRoute().HandlerFunc(notFoundHandler).GetHandler()

	http.ListenAndServe("localhost:3333", r)
}
