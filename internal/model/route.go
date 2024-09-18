package model

// Route struct represents a complete bus route consisting of multiple legs
type Route struct {
	ID   int   `json:"id"`
	Legs []Leg `json:"legs"`
	// Any other relevant fields for the complete route
}
