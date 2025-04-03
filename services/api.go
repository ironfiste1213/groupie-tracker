package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mimo/models"
)

const (
	artistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	locationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	datesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	relationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

// FetchArtists retrieves all artists from the API
func FetchArtists() (models.Artists, error) {
	resp, err := http.Get(artistsURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading artists response: %v", err)
	}

	var artists models.Artists
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, fmt.Errorf("error unmarshaling artists: %v", err)
	}

	return artists, nil
}

func FetchRelationd() (models.Relations, error) {
	resp, err := http.Get(relationsURL)
	if err != nil {
		return models.Relations{}, fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Relations{}, fmt.Errorf("error reading artists response: %v", err)
	}

	var relations models.Relations
	if err := json.Unmarshal(body, &relations); err != nil {
		return models.Relations{}, fmt.Errorf("error unmarshaling artists: %v", err)
	}

	return relations, nil
}
// Additional fetch methods for locations, dates, and relations can be implemented similarly

//func Fetchlocations() (models.Locations, error) {
//	resp, err := http.Get(locationsURL)
//	if err != nil {
//		return nil, fmt.Errorf("error fetching locations: %v", err)
//	}
//	defer resp.Body.Close()
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, fmt.Errorf("error reading locations response: %v", err)
//	}
//
//	var locations  models.Locations
//	if err := json.un
//
