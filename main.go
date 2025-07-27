package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippet box"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	//log message to notify that the server is starting
	log.Print("The server is running on port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
