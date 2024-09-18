// Package model defines the Student struct and related types.
package model

// Student struct represents a student's information
// The `json:` tags specify how these fields should be named in JSON
type Student struct {
	FirstName         string `json:"first_name"`
	MiddleName        string `json:"middle_name,omitempty"` // Optional field
	LastName          string `json:"last_name"`
	AddressID         int    `json:"address_id"`
	ID                int    `json:"student_id"`
	SchoolID          int    `json:"school_id"`
	Grade             string `json:"grade"`
	SpecialNeeds      bool   `json:"special_needs"`         // Indicate if student is special needs
	NeedsNotes        string `json:"needs_notes,omitempty"` // Notes for special needs
	GuardianPrimary   string `json:"guardian_primary"`
	PrimaryPhone      string `json:"primary_phone"`
	GuardianSecondary string `json:"guardian_secondary,omitempty"`
	SecondaryPhone    string `json:"secondary_phone,omitempty"`
	StopID            int    `json:"stop_id"`
	AMStop            bool   `json:"am_stop"`
	PMStop            bool   `json:"pm_stop"`
}
