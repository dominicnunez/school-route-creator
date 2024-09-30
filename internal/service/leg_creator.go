package service

import (
	"fmt"
	"route-creator/internal/model"
)

func separateStudents(students []model.Student) (gened []model.Student, sped []model.Student) {
	for _, student := range students {
		if student.SpecialNeeds {
			sped = append(sped, student)
		} else {
			gened = append(gened, student)
		}
	}
	return gened, sped
}

// CreateSchoolLegs generates route segments based on the provided students, school, and shift
// It takes a school, a slice of students, and a shift as input
// It returns a slice of Leg structs
func CreateSchoolLegs(school model.School, students []model.Student, shift model.Shift) ([]model.Leg, error) {
	if len(students) == 0 {
		return nil, fmt.Errorf("no students provided")
	}

	gened, sped := separateStudents(students)

	// Create legs based on shift
	switch shift {
	case model.AM:
		// TODO: create AM logic
		// Create legs for Gen Ed and Special Needs students
	case model.PM:
		// TODO: create PM logic

	}

	legs := make([]model.Leg, 0)
	return legs, nil
}

func CreateAllSchoolLegs(schools []model.School, students []model.Student, shift model.Shift) ([]model.Leg, error) {
	allLegs := make([]model.Leg, 0)
	for _, school := range schools {
		legs, err := CreateSchoolLegs(school, students, shift)
		if err != nil {
			return nil, err
		}
		allLegs = append(allLegs, legs...)
	}
	return allLegs, nil
}
