package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	port := "8080"
	filepathRoot := "."

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}
