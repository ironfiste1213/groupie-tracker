package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"mimo/models"
	"mimo/services"
	tools "mimo/utils"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle root path
	if r.URL.Path != "/" {
		sendErrorPage(w, "Oops! Page not found", http.StatusNotFound)
		return
	}
	// Fetch artists
	artists, err := services.FetchArtists(services.ArtistsURL)
	if err != nil {
		sendErrorPage(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	// Parse the template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		sendErrorPage(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	// Execute the template with artists data
	err = tmpl.Execute(w, artists)
	if err != nil {
		sendErrorPage(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func ArtistProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Extract artist ID from URL
	// Expected URL format: /artist/1
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	_, err := strconv.Atoi(idStr)
	if err != nil {
		sendErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch all enriched artists
	artist, err := services.FetchArtist(services.ArtistsURL + "/" + idStr)
	relation, err1 := services.FetchRelation(services.RelationsURL + "/" + idStr)
	if err != nil {
		sendErrorPage(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	if err1 != nil {
		sendErrorPage(w, "Failed to fetch artists123", http.StatusInternalServerError)
		return
	}
	if tools.IsZeroArtist(artist) {
		sendErrorPage(w, "Artist not found", http.StatusNotFound)
		return
	}
	
	if tools.IsZeroRelation(relation) {
		sendErrorPage(w, "there is no relation", http.StatusNotFound)
		return
	}
	// Parse and execute template
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		sendErrorPage(w, "Failed to parse template", http.StatusInternalServerError)
		return
	}
	// Define a struct to hold artist and relation
	type ProfileData struct {
		Artist   *models.Artist
		Relation *models.Relation
	}
	// Create an instance of ProfileData
	data := ProfileData{
		Artist:   &artist,
		Relation: &relation,
	}
	err = tmpl.Execute(w, data)

	if err != nil {
		sendErrorPage(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

// Helper: Sends an error page with appropriate HTTP status code.
func sendErrorPage(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		sendErrorPage(w, "Oops! Something went wrong.", http.StatusInternalServerError)
		return
	}
	type PageData struct {
		ErrorMessage string
		Code         int
	}
	data := &PageData{
		ErrorMessage: message,
		Code:         code,
	}
	tmpl.Execute(w, data)
}
