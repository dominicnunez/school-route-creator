// Package data provides an in-memory implementation of the DataStore interface for testing.
package data

import "route-creator/internal/model"

// MemoryStore implements DataStore using in-memory slices
// This is a struct that will hold our data in memory
type MemoryStore struct {
	students []model.Student // Slice of Student structs
	schools  []model.School  // Slice of School structs
}

// NewMemoryStore is a constructor function for MemoryStore
// It initializes and returns a pointer to a new MemoryStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		students: make([]model.Student, 0), // Initialize an empty slice for students
		schools:  make([]model.School, 0),  // Initialize an empty slice for schools
	}
}

// AddStudent adds a new student to the MemoryStore
// The (ms *MemoryStore) syntax defines this as a method on MemoryStore
func (ms *MemoryStore) AddStudent(student model.Student) error {
	ms.students = append(ms.students, student) // Append the new student to the slice
	return nil                                 // Return nil as there's no error
}

// AddSchool adds a new school to the MemoryStore
func (ms *MemoryStore) AddSchool(school model.School) error {
	ms.schools = append(ms.schools, school) // Append the new school to the slice
	return nil
}

// GetStudents returns all students from the MemoryStore
func (ms *MemoryStore) GetStudents() ([]model.Student, error) {
	return ms.students, nil // Return the slice of students
}

// GetSchools returns all schools from the MemoryStore
func (ms *MemoryStore) GetSchools() ([]model.School, error) {
	return ms.schools, nil // Return the slice of schools
}
