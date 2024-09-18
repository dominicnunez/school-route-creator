package service

import "route-creator/internal/model"

// Shift represents whether the leg is for morning (AM) or afternoon (PM) routes
type Shift int

const (
	AM Shift = iota
	PM
)

// CreateLegs generates route segments based on the provided students, school, and shift
// It takes a school, a slice of students, and a shift as input
// It returns a slice of Leg structs
func CreateLegs(school model.School, students []model.Student, shift Shift) []model.Leg {
	legs := make([]model.Leg, 0)

	// Create legs based on shift
	if shift == AM {
		// TODO: create AM logic

	} else {
		// TODO: create PM logic
	}

	return legs
}
