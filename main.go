package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<h1>Get in Touch...</h1>")
	fmt.Fprint(w, "<h3>@jeanrauwers</h3>")
}

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<h1>Hello to Go Web Server</h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404!!!! Page not Found JR</h1>")
}

func main() {
	router := httprouter.New()

	router.GET("/", home)	
	router.GET("/contact", contact)	
	router.NotFound =  http.HandlerFunc(notFoundHandler)


	http.ListenAndServe("localhost:3333", router)
}
