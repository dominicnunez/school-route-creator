// Package model defines the School struct and related types.
package model

// We import the "time" package to work with dates and times
import "time"

// School struct represents a school's information
type School struct {
	Name        string    `json:"name"`         // Name of the school
	ID          int       `json:"school_id"`    // Unique identifier for the school
	AddressID   int       `json:"address_id"`   // Foreign key referencing the address of the school
	StartTime   time.Time `json:"start_time"`   // When school starts (using Go's time.Time type)
	ReleaseTime time.Time `json:"release_time"` // When school ends (using Go's time.Time type)
}
