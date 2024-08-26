package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// home displays a home page.
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Snippetbox")
}

// snippetView displays a specific snippet.
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// snippetCreate displays a form for creating a new snippet.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Display a form for creating a new snippet...")
}

// snippetCreatePost saves a new snippet.
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Save a new snippet...")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
