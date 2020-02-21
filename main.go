package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Get in Touch</h1>")
	fmt.Fprint(w, "<h3>@jeanrauwers</h3>")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello to Go Web Server</h1>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	http.ListenAndServe(":3000", r)
}
