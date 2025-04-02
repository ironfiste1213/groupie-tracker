
package models

// Location represents concert locations for an artist
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// Locations is a slice of Location for easier handling
type Locations []Location