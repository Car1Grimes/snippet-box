package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from Snippet box"))
}

func viewSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 0 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with id = %d", id)
	w.Write([]byte(msg))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", viewSnippet)
	mux.HandleFunc("POST /snippet/create", createSnippet)

	//log message to notify that the server is starting
	log.Print("The server is running on port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
