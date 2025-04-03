package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"mimo/handlers"
	"mimo/services"
)

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Home page handler
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist/",handlers.ArtistProfileHandler)
	// Start the server
	log.Println("Server starting on :8060")
	log.Fatal(http.ListenAndServe(":8060", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle root path
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Fetch artists
	artists, err := services.FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	// Parse the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	// fmt.Printf("%q\n", artists)
	// Execute the template with artists data
	err = tmpl.Execute(w, artists)
	if err != nil {
		fmt.Println("5")
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
