package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"mimo/models"
	"mimo/services"
)

func ArtistProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract artist ID from URL
	// Expected URL format: /artist/1
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	artistID, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch all enriched artists
	artists, err := services.FetchArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	// Find the specific artist
	var artist *models.Artist
	for _, a := range artists {
		if a.ID == artistID {
			artist = &a
			break
		}
	}

	// Artist not found
	if artist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artist)
	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
