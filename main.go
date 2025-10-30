package main // this is the package name

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the riddler"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting a httpserver at port 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
