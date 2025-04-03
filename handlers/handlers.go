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
	relations, err1 := services.FetchRelationd()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	if err1 != nil {
		http.Error(w, "Failed to fetch artists123", http.StatusInternalServerError)
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
	var relation *models.Relation
	for _, a := range relations {
		if a.ID == artistID {
			relation = &a
			break
		}
	}
	// Artist not found
	if artist == nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}
	// Artist not found
	if relation == nil {
		http.Error(w, "Artist not found1", http.StatusNotFound)
		return
	}
	// Parse and execute template
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	// Define a struct to hold artist and relation
	type ProfileData struct {
		Artist   *models.Artist
		Relation *models.Relation
	}
	// Create an instance of ProfileData
	data := ProfileData{
		Artist:   artist,
		Relation: relation,
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
