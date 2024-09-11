// Package model defines the School struct and related types.
package model

// We import the "time" package to work with dates and times
import "time"

// School struct represents a school's information
type School struct {
	Name      string    `json:"name"`       // Name of the school
	Address   string    `json:"address"`    // Address of the school
	StartTime time.Time `json:"start_time"` // When school starts (using Go's time.Time type)
	EndTime   time.Time `json:"end_time"`   // When school ends (using Go's time.Time type)
}
