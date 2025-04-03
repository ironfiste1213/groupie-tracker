package main

import (
	"log"
	"mimo/handlers"
	"net/http"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page handler
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist/", handlers.ArtistProfileHandler)
	// Start the server
	log.Println("Server starting on :8060")
	log.Fatal(http.ListenAndServe(":8060", nil))
}
