package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/snippet/view", viewSnippet)

	log.Print("Starting a server at localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
