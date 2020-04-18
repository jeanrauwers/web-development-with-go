package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Get in Touch</h1>")
	fmt.Fprint(w, "<h3>@jeanrauwers</h3>")
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>404!!!! Page not Found</h1>")
}

func main() {
	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.NotFoundHandler = r.NewRoute().HandlerFunc(notFoundHandler).GetHandler()

	http.ListenAndServe("localhost:3333", r)
}
