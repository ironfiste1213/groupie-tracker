package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"mimo/models"
	"mimo/services"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Handle root path
	if r.URL.Path != "/" {
		sendErrorPage(w, "Oops! Page not found", http.StatusNotFound)
		return
	}

	// Fetch artists
	artists, err := services.FetchArtists()
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
	artistID, err := strconv.Atoi(idStr)
	if err != nil {
		sendErrorPage(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch all enriched artists
	artists, err := services.FetchArtists()
	relations, err1 := services.FetchRelationd()
	if err != nil {
		sendErrorPage(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	if err1 != nil {
		sendErrorPage(w, "Failed to fetch artists123", http.StatusInternalServerError)
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
	for _, a := range relations.Index {
		if a.ID == artistID {
			relation = &a
			break
		}
	}
	// Artist not found
	if artist == nil {
		sendErrorPage(w, "Artist not found", http.StatusNotFound)
		return
	}
	// Artist not found
	if relation == nil {
		sendErrorPage(w, "Artist not found1", http.StatusNotFound)
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
		Artist:   artist,
		Relation: relation,
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
