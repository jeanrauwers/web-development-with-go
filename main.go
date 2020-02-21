package main

import (
	"fmt"
	"net/http"
)

func handleReturn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visited our page")
	fmt.Fprint(w, "<h1>Hello` from go server</h1>")
}

func main() {
	http.HandleFunc("/", handleReturn)
	http.ListenAndServe(":3000", nil)
}
