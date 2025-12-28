package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the riddler creating the snippetbox")) // what this line does I have to be clear in future once understanding the fundamentals of the language
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting the server on localhost:4000")
	err := http.ListenAndServe(":4000", mux) // since this line has potential chances of failing or throwing an error we will have to look for the error.
	log.Fatal(err)
}
