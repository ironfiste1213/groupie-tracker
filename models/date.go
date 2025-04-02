package models

// Date represents concert dates for an artist
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Dates is a slice of Date for easier handling
type Dates []Date