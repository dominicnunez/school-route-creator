// Package main is the entry point of our application
package main

// We import the "time" package to work with dates and times
import "time"

// Student struct represents a student's information
// The `json:` tags specify how these fields should be named in JSON
type Student struct {
	Name         string `json:"name"`          // Student's full name
	Address      string `json:"address"`       // Student's home address
	SpecialNeeds string `json:"special_needs"` // Any special needs or requirements the student has
}

// School struct represents a school's information
type School struct {
	Name      string    `json:"name"`       // Name of the school
	Address   string    `json:"address"`    // Address of the school
	StartTime time.Time `json:"start_time"` // When school starts (using Go's time.Time type)
	EndTime   time.Time `json:"end_time"`   // When school ends (using Go's time.Time type)
}

// Route struct represents a bus route
type Route struct {
	Students []Student `json:"students"` // A slice (dynamic array) of Student structs
	School   School    `json:"school"`   // The School this route is for
}

// These variables will store our data
// In a real application, we'd probably use a database instead
var students []Student // A slice to hold all our students
var schools []School   // A slice to hold all our schools
