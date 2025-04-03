package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"mimo/models"
)

const (
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationsURL = "https://groupietrackers.herokuapp.com/api/locations"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
	RelationsURL = "https://groupietrackers.herokuapp.com/api/relation"
)

// FetchArtists retrieves all artists from the API
func FetchArtists(URL string) (models.Artists, error) {
	resp, err := http.Get(URL)
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

func FetchArtist(URL string) (models.Artist, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return models.Artist{}, fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Artist{}, fmt.Errorf("error reading artists response: %v", err)
	}

	var artist models.Artist
	if err := json.Unmarshal(body, &artist); err != nil {
		return models.Artist{}, fmt.Errorf("error unmarshaling artists: %v", err)
	}

	return artist, nil
}

func FetchRelation(URL string) (models.Relation, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return models.Relation{}, fmt.Errorf("error fetching artists: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.Relation{}, fmt.Errorf("error reading artists response: %v", err)
	}

	var relation models.Relation
	if err := json.Unmarshal(body, &relation); err != nil {
		return models.Relation{}, fmt.Errorf("error unmarshaling artists: %v", err)
	}

	return relation, nil
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
