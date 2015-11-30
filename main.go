package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	// Initilize db
	SetupDB()

	// Declare routes
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/game/{id:[0-9]+}", gameDetailHandler)
	http.Handle("/", r)

	// For static contents
	http.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Serve
	http.ListenAndServe(":8080", nil)
}
