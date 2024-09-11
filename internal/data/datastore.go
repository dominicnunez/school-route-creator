// Package data defines the DataStore interface and provides
// implementations for different types of data storage.
package data

import "route-creator/internal/model"

// DataStore interface defines methods for data operations
// Interfaces in Go provide a way to specify the behavior of an object
type DataStore interface {
	AddStudent(student model.Student) error
	AddSchool(school model.School) error
	GetStudents() ([]model.Student, error)
	GetSchools() ([]model.School, error)
}
