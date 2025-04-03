package tools

import "mimo/models"

// Helper: Checks if an artist struct is empty (zero value).
func IsZeroArtist(a models.Artist) bool {
	return a.Name == "" && a.ID == 0 // Add other fields if necessary
}
// Helper: Checks if an relation struct is empty (zero value).
func IsZeroRelation(a models.Relation) bool {
	return a.DatesLocations == nil && a.ID == 0
}