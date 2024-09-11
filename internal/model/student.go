// Package model defines the Student struct and related types.
package model

// Student struct represents a student's information
// The `json:` tags specify how these fields should be named in JSON
type Student struct {
	Name         string `json:"name"`          // Student's full name
	Address      string `json:"address"`       // Student's home address
	SpecialNeeds string `json:"special_needs"` // Any special needs or requirements the student has
}
