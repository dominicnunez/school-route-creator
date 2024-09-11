// Package model defines the Route struct and related types.
package model

// Route struct represents a bus route
type Route struct {
	Students []Student `json:"students"` // A slice (dynamic array) of Student structs
	School   School    `json:"school"`   // The School this route is for
}
