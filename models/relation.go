package models

// Relation represents the relationship between artists and their concert dates/locations
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}