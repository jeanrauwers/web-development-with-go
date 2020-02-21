package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleReturn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Helloss` from go server</h1>")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", handleReturn)
	http.ListenAndServe(":3000", r)
}
